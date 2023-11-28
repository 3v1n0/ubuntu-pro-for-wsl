package endtoend_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/canonical/ubuntu-pro-for-windows/common"
	"github.com/canonical/ubuntu-pro-for-windows/common/wsltestutils"
	"github.com/stretchr/testify/require"
	"github.com/ubuntu/gowsl"
)

func testSetup(t *testing.T) {
	t.Helper()
	ctx := context.Background()

	err := gowsl.Shutdown(ctx)
	require.NoError(t, err, "Setup: could not shut WSL down")

	err = stopAgent(ctx)
	require.NoError(t, err, "Setup: could not stop the agent")

	err = assertCleanRegistry()
	require.NoError(t, err, "Setup: registry is polluted, potentially by a previous test")

	err = assertCleanFilesystem()
	require.NoError(t, err, "Setup: local app data is polluted, potentially by a previous test")

	t.Cleanup(func() {
		err := errors.Join(
			stopAgent(ctx),
			cleanupRegistry(),
			cleanupFilesystem(),
		)
		// Cannot assert: the test is finished already
		if err != nil {
			log.Printf("Cleanup: %v", err)
		}
	})
}

//nolint:revive // testing.T must precede the context
func registerFromTestImage(t *testing.T, ctx context.Context) string {
	t.Helper()

	distroName := wsltestutils.RandomDistroName(t)
	t.Logf("Registering distro %q", distroName)
	defer t.Logf("Registered distro %q", distroName)

	_ = wsltestutils.PowershellImportDistro(t, ctx, distroName, testImagePath)
	return distroName
}

// startAgent starts the GUI (without interacting with it) and waits for the Agent to start.
// A single command line argument is expected. Additionally, environment variable overrides
// in a form of "key=value" strings can be appended to the current environment.
// It stops the agent upon cleanup. If the cleanup fails, the testing will be stopped.
//
//nolint:revive // testing.T must precede the contex
func startAgent(t *testing.T, ctx context.Context, arg string, environ ...string) (cleanup func()) {
	t.Helper()

	t.Log("Starting agent")

	out, err := powershellf(ctx, "(Get-AppxPackage CanonicalGroupLimited.UbuntuProForWindows).InstallLocation").CombinedOutput()
	require.NoError(t, err, "could not locate ubuntupro.exe: %v. %s", err, out)

	ubuntupro := filepath.Join(strings.TrimSpace(string(out)), "gui", "ubuntupro.exe")
	//nolint:gosec // The executable is located at the Appx directory
	cmd := exec.CommandContext(ctx, ubuntupro, arg)

	var buff bytes.Buffer
	cmd.Stdout = &buff
	cmd.Stderr = &buff

	if environ != nil {
		cmd.Env = append(cmd.Environ(), environ...)
	}

	err = cmd.Start()
	require.NoError(t, err, "Setup: could not start agent")

	cleanup = func() {
		t.Log("Cleanup: stopping agent process")

		if err := stopAgent(ctx); err != nil {
			// We have to abort because the tests become coupled via the agent
			log.Fatalf("Could not kill ubuntu-pro-agent process: %v: %s", err, out)
		}

		//nolint:errcheck // Nothing we can do about it
		cmd.Process.Kill()

		//nolint:errcheck // We know that the previous "Kill" stopped it
		cmd.Wait()
		t.Logf("Agent stopped. Stdout+stderr: %s", buff.String())
	}

	defer func() {
		if t.Failed() {
			cleanup()
		}
	}()

	home := os.Getenv("UserProfile")
	require.NotEmptyf(t, home, "Agent setup: $env:UserProfile should not be empty")

	require.Eventually(t, func() bool {
		_, err := os.Stat(filepath.Join(home, ".ubuntupro"))
		if errors.Is(err, fs.ErrNotExist) {
			return false
		}
		if err != nil {
			t.Logf("Agent setup: could not read address file: %v", err)
			return false
		}
		return true
	}, 30*time.Second, 100*time.Millisecond, "Agent never started serving")

	t.Log("Started agent")
	return cleanup
}

// stopAgent kills the process for the Windows Agent.
func stopAgent(ctx context.Context) error {
	const process = "ubuntu-pro-agent"

	out, err := powershellf(ctx, "Stop-Process -ProcessName %q", process).CombinedOutput()
	if err == nil {
		return nil
	}

	if strings.Contains(string(out), "NoProcessFoundForGivenName") {
		return nil
	}

	return fmt.Errorf("could not stop process %q: %v. %s", process, err, out)
}

//nolint:revive // testing.T must precede the context
func distroIsProAttached(t *testing.T, ctx context.Context, d gowsl.Distro) (bool, error) {
	t.Helper()

	var stdout, stderr bytes.Buffer

	cmd := d.Command(ctx, "pro status --format=json")

	// We need separate Stdout and Stderr. We cannot combine them because the
	// pro client prints warnings to Stderr, which makes the combined output
	// invalid JSON. We also cannot ignore either:
	// - We need StdOut to parse the JSON output.
	// - We need Stderr to read the error message in case the command fails.
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return false, fmt.Errorf("could not call pro status: %v.\nSTDOUT: %s\nSTDERR: %s", err, &stdout, &stderr)
	}

	var response struct {
		Attached bool
	}
	out := stdout.Bytes()
	if err := json.Unmarshal(out, &response); err != nil {
		return false, fmt.Errorf("could not parse pro status response: %v: %s", err, out)
	}

	return response.Attached, nil
}

//nolint:revive // testing.T must precede the context
func logWslProServiceOnError(t *testing.T, ctx context.Context, d gowsl.Distro) {
	t.Helper()

	if !t.Failed() {
		return
	}

	out, err := d.Command(ctx, "journalctl -b --no-pager -u wsl-pro.service").CombinedOutput()
	if err != nil {
		t.Logf("could not access WSL Pro Service logs: %v\n%s\n", err, out)
		return
	}
	t.Logf("WSL Pro Service logs:\n%s\n", out)
}

//nolint:revive // testing.T must precede the context
func logWindowsAgentOnError(t *testing.T, ctx context.Context) {
	t.Helper()

	if !t.Failed() {
		return
	}

	localAppData := os.Getenv("LocalAppData")
	if localAppData == "" {
		t.Log("could not find Windows Agent's logs: $env:LocalAppData is not assigned")
		return
	}

	// The virtualized LocalAppData is located under a path similar to this one:
	//
	//    %LocalAppData%/Packages/CanonicalGroupLimited.UbuntuProForWindows_hhj52ngek5ykr/LocalCache/Local
	//                                                                      ^~~~~~~~~~~~~
	//                                                                      This part changes from version to version
	//
	packageDir := filepath.Join(localAppData, "Packages", "CanonicalGroupLimited.UbuntuProForWindows_*")

	out, err := powershellf(ctx, "(Get-Item %q).FullName", packageDir).CombinedOutput()
	if err != nil {
		t.Logf("could not find Windows Agent's logs: could not find virtualized LocalAppData: %v", err)
		return
	}

	// Remove trailing endline
	packageDir = strings.TrimSpace(string(out))

	logsPath := filepath.Join(packageDir, "LocalCache", "Local", common.LocalAppDataDir, ".ubuntupro.log")
	out, err = os.ReadFile(logsPath)
	if err != nil {
		t.Logf("could not read Windows Agent's logs at %q: %v", logsPath, err)
		return
	}

	t.Logf("Windows Agent's logs:\n%s\n", out)
}
