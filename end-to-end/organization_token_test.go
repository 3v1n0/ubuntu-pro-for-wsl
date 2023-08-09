package endtoend_test

import (
	"context"
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	wsl "github.com/ubuntu/gowsl"
	"golang.org/x/sys/windows/registry"
)

func TestOrganizationProvidedToken(t *testing.T) {
	type whenToken int
	const (
		never whenToken = iota
		beforeDistroRegistration
		afterDistroRegistration
	)

	testCases := map[string]struct {
		whenToken whenToken

		wantAttached bool
	}{
		"Success when the subscription is active before registration":   {whenToken: beforeDistroRegistration, wantAttached: true},
		"Success when the subscription is activated after registration": {whenToken: afterDistroRegistration, wantAttached: true},

		"Error when there is no active subscription": {whenToken: never, wantAttached: false},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()

			testSetup(t)

			if tc.whenToken == beforeDistroRegistration {
				activateOrgSubscription(t)
				startAgent(t, ctx)
			}

			// Distro setup
			name := registerFromTestImage(t, ctx)
			d := wsl.NewDistro(ctx, name)

			defer logWslProServiceJournal(t, ctx, d)

			out, err := d.Command(ctx, "exit 0").CombinedOutput()
			require.NoErrorf(t, err, "Setup: could not wake distro up: %v. %s", err, out)

			if tc.whenToken == afterDistroRegistration {
				err := d.Terminate()
				require.NoError(t, err, "could not restart distro")

				activateOrgSubscription(t)
				startAgent(t, ctx)

				out, err := d.Command(ctx, "exit 0").CombinedOutput()
				require.NoErrorf(t, err, "Setup: could not wake distro up: %v. %s", err, out)
			}

			// Give the agent some time to pro-attach
			time.Sleep(5 * time.Second)

			// Validate that the distro was attached
			out, err = d.Command(ctx, "pro status --format=json").Output()
			require.NoErrorf(t, err, "Setup: could not call pro status: %v. %s", err, out)

			var response struct {
				Attached bool
			}
			err = json.Unmarshal(out, &response)
			require.NoError(t, err, "could not parse pro status response: %s", out)

			if !tc.wantAttached {
				require.False(t, response.Attached, "distro should not have been Pro attached")
				return
			}

			require.True(t, response.Attached, "distro should have been Pro attached")
		})
	}
}

func activateOrgSubscription(t *testing.T) {
	t.Helper()

	token := os.Getenv(proTokenEnv)
	require.NotEmptyf(t, token, "Setup: environment variable %q should contain a valid pro token, but is empty", proTokenEnv)

	key, _, err := registry.CreateKey(registry.CURRENT_USER, registryPath, registry.WRITE)
	require.NoErrorf(t, err, "Setup: could not open UbuntuPro registry key")
	defer key.Close()

	err = key.SetStringValue("ProTokenOrg", token)
	require.NoError(t, err, "could not write token in registry")
}

//nolint:revive // testing.T must precede the context
func logWslProServiceJournal(t *testing.T, ctx context.Context, d wsl.Distro) {
	t.Helper()

	out, err := d.Command(ctx, "journalctl --no-pager -u wsl-pro.service").CombinedOutput()
	if err != nil {
		t.Logf("could not access logs: %v\n%s\n", err, out)
	}
	t.Logf("wsl-pro-service logs:\n%s\n", out)
}