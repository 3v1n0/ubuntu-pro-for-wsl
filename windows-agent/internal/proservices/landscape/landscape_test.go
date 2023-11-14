package landscape_test

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	landscapeapi "github.com/canonical/landscape-hostagent-api"
	"github.com/canonical/ubuntu-pro-for-windows/common/golden"
	"github.com/canonical/ubuntu-pro-for-windows/common/wsltestutils"
	"github.com/canonical/ubuntu-pro-for-windows/mocks/landscape/landscapemockservice"
	"github.com/canonical/ubuntu-pro-for-windows/windows-agent/internal/config"
	"github.com/canonical/ubuntu-pro-for-windows/windows-agent/internal/distros/database"
	"github.com/canonical/ubuntu-pro-for-windows/windows-agent/internal/distros/distro"
	"github.com/canonical/ubuntu-pro-for-windows/windows-agent/internal/distros/task"
	"github.com/canonical/ubuntu-pro-for-windows/windows-agent/internal/proservices/landscape"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	wsl "github.com/ubuntu/gowsl"
	wslmock "github.com/ubuntu/gowsl/mock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func TestMain(m *testing.M) {
	log.SetLevel(log.DebugLevel)

	exit := m.Run()
	defer os.Exit(exit)
}

func TestConnect(t *testing.T) {
	if wsl.MockAvailable() {
		t.Parallel()
	}

	certPath := filepath.Join(golden.TestFamilyPath(t), "certificates")
	t.Cleanup(func() {
		if err := os.RemoveAll(certPath); err != nil {
			log.Printf("Failed to remove certificates directory: %v", err)
		}
	})

	err := os.MkdirAll(certPath, 0600)
	require.NoError(t, err, "Setup: could not create certificates directory")

	err = generateTempCertificate(certPath)
	require.NoError(t, err, "Setup: could not generate certificates")

	err = os.WriteFile(filepath.Join(certPath, "bad-certificate.pem"), []byte("This is not a valid certificate."), 0600)
	require.NoError(t, err, "Setup: could not create bad certificate")

	testCases := map[string]struct {
		precancelContext   bool
		serverNotAvailable bool

		landscapeUIDReadErr  bool
		landscapeUIDWriteErr bool

		noLandscapeURL bool
		tokenErr       bool

		requireCertificate         bool
		breakLandscapeClientConfig bool

		breakUIDFile bool
		uid          string

		wantErr           bool
		wantDistroSkipped bool
	}{
		"Success in first contact":        {},
		"Success in non-first contact":    {uid: "123"},
		"Success with non-empty config":   {},
		"Success with an SSL certificate": {requireCertificate: true},

		"Error when the context is cancelled before Connected": {precancelContext: true, wantErr: true},
		"Error when the landscape URL cannot be retrieved":     {noLandscapeURL: true, wantErr: true},
		"Error when the landscape UID cannot be retrieved":     {landscapeUIDReadErr: true, wantErr: true},
		"Error when the landscape UID cannot be stored":        {landscapeUIDWriteErr: true, wantErr: true},
		"Error when the server cannot be reached":              {serverNotAvailable: true, wantErr: true},
		"Error when the first-contact SendUpdatedInfo fails":   {tokenErr: true, wantErr: true},
		"Error when the config cannot be accessed":             {breakLandscapeClientConfig: true, wantErr: true},
		"Error when the config cannot be parsed":               {wantErr: true},
		"Error when the SSL certificate cannot be read":        {wantErr: true},
		"Error when the SSL certificate is not valid":          {wantErr: true},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			if wsl.MockAvailable() {
				t.Parallel()
				ctx = wsl.WithMock(ctx, wslmock.New())
			}

			lis, server, mockService := setUpLandscapeMock(t, ctx, "localhost:", tc.requireCertificate)
			defer lis.Close()

			conf := &mockConfig{
				proToken: "TOKEN",

				// We trigger an error on first-contact SendUpdatedInfo by erroring out in conf.ProToken()
				proTokenErr: tc.tokenErr,

				// We trigger errors trying to read or write to/from the registry
				landscapeUIDErr:    tc.landscapeUIDReadErr,
				setLandscapeUIDErr: tc.landscapeUIDWriteErr,

				// We trigger an error when deciding to use a certificate or not
				landscapeConfigErr: tc.breakLandscapeClientConfig,

				landscapeAgentUID: tc.uid,
			}

			out, err := os.ReadFile(filepath.Join(golden.TestFixturePath(t), "landscape.conf"))
			if errors.Is(err, os.ErrNotExist) {
				// This fixture is not compulsory
				out = []byte{}
				err = nil
			}
			require.NoError(t, err, "Setup: could not load landscape config")

			conf.landscapeClientConfig = string(out)
			if !tc.noLandscapeURL {
				conf.landscapeClientConfig = fmt.Sprintf("[host]\nurl=%q\n\n%s", lis.Addr(), conf.landscapeClientConfig)
			}

			if !tc.serverNotAvailable {
				//nolint:errcheck // We don't care about these errors
				go server.Serve(lis)
				defer server.Stop()
			}

			db, err := database.New(ctx, t.TempDir(), conf)
			require.NoError(t, err, "Setup: database New should not return an error")

			distroName, _ := wsltestutils.RegisterDistro(t, ctx, true)
			_, err = db.GetDistroAndUpdateProperties(ctx, distroName, distro.Properties{})
			require.NoError(t, err, "Setup: GetDistroAndUpdateProperties should return no errors")

			client, err := landscape.NewClient(conf, db)
			require.NoError(t, err, "Setup: NewClient should return no errrors")

			ctx, cancel := context.WithCancel(ctx)
			defer cancel()

			if tc.precancelContext {
				cancel()
			}

			err = client.Connect(ctx)
			if tc.wantErr {
				require.Error(t, err, "Connect should return an error")
				require.False(t, client.Connected(), "Connected should have returned false after failing to connect")
				return
			}
			require.NoError(t, err, "Connect should return no errors")
			defer client.Stop(ctx)

			require.True(t, client.Connected(), "Connected should have returned false after succeeding to connect")

			require.Eventually(t, func() bool {
				return len(mockService.MessageLog()) > 0
			}, 10*time.Second, 100*time.Millisecond, "Landscape server should receive a message from the client")

			client.Stop(ctx)
			require.NotPanics(t, func() { client.Stop(ctx) }, "client.Stop should not panic, even when called twice")

			require.False(t, client.Connected(), "Connected should have returned false after disconnecting")

			wantUID := tc.uid
			if tc.uid == "" {
				wantUID = "ServerAssignedUID"
			}
			requireHasPrefix(t, wantUID, conf.landscapeAgentUID, "Landscape client UID was not set properly")

			server.Stop()
			lis.Close()

			messages := mockService.MessageLog()
			require.Len(t, messages, 1, "Exactly one message should've been sent to Landscape")
		})
	}
}

// generateTempCertificate generates a self-signed certificate valid for one hour. Both the
// certificate and the private key are stored in the specified path.
func generateTempCertificate(path string) error {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return fmt.Errorf("could not generate keys: %v", err)
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName:   "CanonicalGroupLimited",
			Country:      []string{"US"},
			Organization: []string{"Canonical"},
		},
		IPAddresses: []net.IP{net.ParseIP("127.0.0.1")},
		NotBefore:   time.Now(),
		NotAfter:    time.Now().Add(time.Hour),
	}

	cert, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		return fmt.Errorf("could not create certificate: %s", err)
	}

	// Marshal and write certificate
	out := &bytes.Buffer{}
	if err := pem.Encode(out, &pem.Block{Type: "CERTIFICATE", Bytes: cert}); err != nil {
		return fmt.Errorf("could not encode certificate: %v", err)
	}

	if err := os.WriteFile(filepath.Join(path, "cert.pem"), out.Bytes(), 0600); err != nil {
		return fmt.Errorf("could not write certificate to file: %v", err)
	}

	// Marshal and write private key
	key, err := x509.MarshalECPrivateKey(priv)
	if err != nil {
		return fmt.Errorf("could not marshal private key: %v", err)
	}

	out = &bytes.Buffer{}
	if err := pem.Encode(out, &pem.Block{Type: "EC PRIVATE KEY", Bytes: key}); err != nil {
		return fmt.Errorf("could not encode private key: %v", err)
	}

	if err := os.WriteFile(filepath.Join(path, "key.pem"), out.Bytes(), 0600); err != nil {
		return fmt.Errorf("could not write private key to file: %v", err)
	}

	return nil
}

func TestSendUpdatedInfo(t *testing.T) {
	if wsl.MockAvailable() {
		t.Parallel()
	}

	testCases := map[string]struct {
		tokenErr bool
		stateErr bool

		precancelContext     bool
		disconnectBeforeSend bool
		distroIsRunning      bool
		distroIsUnregistered bool

		wantErr           bool
		wantDistroSkipped bool
	}{
		"Success with a stopped distro":                     {},
		"Success with a running distro":                     {distroIsRunning: true},
		"Success when the distro State cannot be retreived": {stateErr: true, wantDistroSkipped: true},

		"Error when the token cannot be retreived":                           {tokenErr: true, wantErr: true},
		"Error when attempting to SendUpdatedInfo after having disconnected": {disconnectBeforeSend: true, wantErr: true},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			if wsl.MockAvailable() {
				t.Parallel()
				mock := wslmock.New()
				mock.StateError = tc.stateErr
				ctx = wsl.WithMock(ctx, mock)
			}

			lis, server, mockService := setUpLandscapeMock(t, ctx, "localhost:", false)

			conf := &mockConfig{
				proToken:              "TOKEN",
				landscapeClientConfig: fmt.Sprintf("[host]\nurl=%q\n", lis.Addr()),
			}

			//nolint:errcheck // We don't care about these errors
			go server.Serve(lis)
			defer server.Stop()

			db, err := database.New(ctx, t.TempDir(), conf)
			require.NoError(t, err, "Setup: database New should not return an error")

			distroName, _ := wsltestutils.RegisterDistro(t, ctx, true)
			props := distro.Properties{
				DistroID:    "Cool Ubuntu",
				VersionID:   "NewerThanYours",
				PrettyName:  "😎 Cool guy 🎸",
				Hostname:    "CoolMachine",
				ProAttached: true,
			}

			distro, err := db.GetDistroAndUpdateProperties(ctx, distroName, props)
			require.NoError(t, err, "Setup: GetDistroAndUpdateProperties should return no errors")

			const hostname = "HOSTNAME"

			client, err := landscape.NewClient(conf, db, landscape.WithHostname(hostname))
			require.NoError(t, err, "Landscape NewClient should not return an error")

			if tc.distroIsRunning {
				err := distro.LockAwake()
				//nolint:errcheck // Nothing we can do about it
				defer distro.ReleaseAwake()
				require.NoError(t, err, "Setup: could not keep distro alive")
			} else {
				d := wsl.NewDistro(ctx, distroName)
				err := d.Terminate()
				require.NoError(t, err, "Setup: could not terminate the distro")
			}

			err = client.Connect(ctx)
			require.NoError(t, err, "Setup: Connect should return no errors")
			defer client.Stop(ctx)

			// Defining wants
			wantUIDprefix := "ServerAssignedUID"
			wantHostname := hostname
			wantHostToken := conf.proToken
			wantDistroID := distroName
			wantDistroName := props.Hostname
			wantDistroVersionID := props.VersionID
			wantDistroState := landscapeapi.InstanceState_Stopped
			if tc.distroIsRunning {
				wantDistroState = landscapeapi.InstanceState_Running
			}

			// Asserting on the first-contact SendUpdatedInfo
			require.Eventually(t, func() bool {
				return len(mockService.MessageLog()) > 0
			}, 10*time.Second, 100*time.Millisecond, "Landscape server should receive a message from the client")

			messages := mockService.MessageLog()
			require.Len(t, messages, 1, "Exactly one message should've been sent to Landscape")
			msg := &messages[0] // Pointer to avoid copying mutex

			assert.Empty(t, msg.UID, "First UID received by the server should be empty")
			assert.Equal(t, wantHostname, msg.Hostname, "Mismatch between local host ID and that received by the server")
			assert.Equal(t, wantHostToken, msg.Token, "Mismatch between local host pro token and those received by the server")

			if tc.wantDistroSkipped {
				require.Empty(t, msg.Instances, "No distro should've been sent to Landscape")
			} else {
				require.Len(t, msg.Instances, 1, "Exactly one distro should've been sent to Landscape")
				got := msg.Instances[0]
				assert.Equal(t, wantDistroID, got.ID, "Mismatch between local distro Id and that received by the server")
				assert.Equal(t, wantDistroName, got.Name, "Mismatch between local distro Name and that received by the server")
				assert.Equal(t, wantDistroVersionID, got.VersionID, "Mismatch between local distro VersionId and that received by the server")
				assert.Equal(t, wantDistroState, got.InstanceState, "Mismatch between local distro InstanceState and that received by the server")
			}

			// Exiting if previous assert battery failed
			if t.Failed() {
				t.FailNow()
			}

			// Setting up SendUpdatedInfo
			conf.proTokenErr = tc.tokenErr
			conf.proToken = "NEW_TOKEN"

			if tc.disconnectBeforeSend {
				client.Stop(ctx)
			}

			wantHostToken = conf.proToken

			if !tc.distroIsRunning {
				d := wsl.NewDistro(ctx, distroName)
				err := d.Terminate()
				require.NoError(t, err, "Setup: could not terminate distro")
			}

			err = client.SendUpdatedInfo(ctx)
			if tc.wantErr {
				require.Error(t, err, "SendUpdatedInfo should have returned an error")
				return
			}
			require.NoError(t, err, "SendUpdatedInfo should send no error")

			// Asserting on the second SendUpdatedInfo
			require.Eventually(t, func() bool {
				return len(mockService.MessageLog()) > 1
			}, 10*time.Second, 100*time.Millisecond, "Landscape server should receive a second message from the client")

			messages = mockService.MessageLog()
			require.Len(t, messages, 2, "Exactly two messages should've been sent to Landscape")
			msg = &messages[1] // Pointer to avoid copying mutex

			assertHasPrefix(t, wantUIDprefix, msg.UID, "Mismatch between local host ID and that received by the server")
			assert.Equal(t, wantHostname, msg.Hostname, "Mismatch between local host hostname and that received by the server")
			assert.Equal(t, wantHostToken, msg.Token, "Mismatch between local host pro token and those received by the server")
			if tc.wantDistroSkipped {
				require.Empty(t, msg.Instances, "No distro should've been sent to Landscape")
			} else {
				require.Len(t, msg.Instances, 1, "Exactly one distro should've been sent to Landscape")
				got := msg.Instances[0]
				assert.Equal(t, wantDistroID, got.ID, "Mismatch between local distro Id and that received by the server")
				assert.Equal(t, wantDistroName, got.Name, "Mismatch between local distro Name and that received by the server")
				assert.Equal(t, wantDistroVersionID, got.VersionID, "Mismatch between local distro VersionId and that received by the server")
				assert.Equal(t, wantDistroState, got.InstanceState, "Mismatch between local distro InstanceState and that received by the server ")
			}
		})
	}
}

func requireHasPrefix(t *testing.T, wantPrefix, got string, msgAndArgs ...interface{}) {
	t.Helper()

	if assertHasPrefix(t, wantPrefix, got, msgAndArgs...) {
		return
	}

	t.FailNow()
}

func assertHasPrefix(t *testing.T, wantPrefix, got string, msgAndArgs ...interface{}) bool {
	t.Helper()

	if strings.HasPrefix(got, wantPrefix) {
		return true
	}

	errMsg := fmt.Sprintf("String does not have prefix.\n    Prefix: %s\n    String: %s\n", wantPrefix, got)
	assert.Fail(t, errMsg, msgAndArgs)
	return false
}

func TestAutoReconnection(t *testing.T) {
	ctx := context.Background()
	if wsl.MockAvailable() {
		t.Parallel()
		mock := wslmock.New()
		ctx = wsl.WithMock(ctx, mock)
	}

	lis, server, mockService := setUpLandscapeMock(t, ctx, "localhost:", false)
	defer lis.Close()
	defer server.Stop()

	conf := &mockConfig{
		proToken:              "TOKEN",
		landscapeClientConfig: fmt.Sprintf("[host]\nurl=%q\n", lis.Addr()),
	}

	db, err := database.New(ctx, t.TempDir(), conf)
	require.NoError(t, err, "Setup: database New should not return an error")

	const hostname = "HOSTNAME"

	client, err := landscape.NewClient(conf, db, landscape.WithHostname(hostname))
	require.NoError(t, err, "Landscape NewClient should not return an error")
	defer client.Stop(ctx)

	err = client.Connect(ctx)
	require.Error(t, err, "Connect should have failed because the server is not running")

	require.False(t, client.Connected(), "Client should not be connected because the server is not running")

	//nolint:errcheck // We don't care about these errors
	go server.Serve(lis)
	defer server.Stop()

	require.Eventually(t, client.Connected, 5*time.Second, 500*time.Millisecond, "Client should have reconnected after starting the server")

	hosts := mockService.Hosts()
	require.Len(t, hosts, 1, "Only one client should have connected to the Landscape server")

	for uid := range hosts {
		err = mockService.Disconnect(uid)
		require.NoError(t, err, "Server should not fail to disconnect the client")
	}

	// Fast tickrate to ensure we detect the disconnection. The client reconnects after 1s so we should always win.
	const tick = 10 * time.Millisecond
	require.Eventually(t, func() bool {
		return !client.Connected()
	}, 5*time.Second, tick, "Client should have disconnected after the stream is dropped")

	// Detecting reconnection
	require.Eventually(t, client.Connected, 5*time.Second, 100*time.Millisecond, "Client should have reconnected after the stream is dropped")

	server.Stop()
	require.Eventually(t, func() bool {
		return !client.Connected()
	}, 5*time.Second, 100*time.Millisecond, "Client should have disconnected after the server is stopped")

	// Restart server at the same address
	lis, server, _ = setUpLandscapeMock(t, ctx, lis.Addr().String(), false)
	defer lis.Close()

	//nolint:errcheck // We don't care
	go server.Serve(lis)
	defer server.Stop()

	require.Eventually(t, client.Connected, 5*time.Second, 500*time.Millisecond, "Client should have reconnected after restarting the server")
}

type command int

const (
	cmdAssignHost command = iota
	cmdStart
	cmdStop
	cmdInstall
	cmdUninstall
	cmdSetDefault
	cmdShutdownHost
)

func TestReceiveCommands(t *testing.T) {
	if wsl.MockAvailable() {
		t.Parallel()
	}

	testCases := map[string]struct {
		command command
		mockErr bool

		wantFailure bool
	}{
		"Success receiving a AssignHost command": {command: cmdAssignHost},

		"Success receiving a Start command": {command: cmdStart},
		"Error receiving a Start command":   {command: cmdStart, mockErr: true, wantFailure: true},

		"Success receiving a Stop command": {command: cmdStop},
		"Error receiving a Stop command":   {command: cmdStop, mockErr: true, wantFailure: true},

		"Success receiving a Install command": {command: cmdInstall},
		"Error receiving a Install command":   {command: cmdInstall, mockErr: true, wantFailure: true},

		"Success receiving a Uninstall command": {command: cmdUninstall},
		"Error receiving a Uninstall command":   {command: cmdUninstall, mockErr: true, wantFailure: true},

		"Success receiving a SetDefault command": {command: cmdSetDefault},
		"Error receiving a SetDefault command":   {command: cmdSetDefault, mockErr: true, wantFailure: true},

		"Success receiving a ShutdownHost command": {command: cmdShutdownHost},
		"Error receiving a ShutdownHost command":   {command: cmdShutdownHost, mockErr: true, wantFailure: true},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()

			enableMockErrors := func() {}
			disableMockErrors := func() {}

			if wsl.MockAvailable() {
				t.Parallel()
				mock := wslmock.New()
				ctx = wsl.WithMock(ctx, mock)
				if tc.mockErr {
					enableMockErrors = func() {
						mock.WslLaunchInteractiveError = true      // Breaks start
						mock.InstallError = true                   // Breaks install
						mock.WslUnregisterDistributionError = true // Breaks uninstall
						mock.SetAsDefaultError = true              // Breaks SetDefault
						mock.ShutdownError = true                  // Breaks shutdown
					}
					disableMockErrors = mock.ResetErrors
				}
				defer mock.ResetErrors()
			} else if tc.mockErr {
				t.Skip("This test can only run with the mock")
			}

			lis, server, service := setUpLandscapeMock(t, ctx, "localhost:", false)
			defer lis.Close()

			//nolint:errcheck // We don't care about these errors
			go server.Serve(lis)
			defer server.Stop()

			conf := &mockConfig{
				proToken:              "TOKEN",
				landscapeClientConfig: fmt.Sprintf("[host]\nurl=%q\n", lis.Addr()),
			}

			db, err := database.New(ctx, t.TempDir(), conf)
			require.NoError(t, err, "Setup: database New should not return an error")

			var d *distro.Distro
			if tc.command != cmdInstall {
				distroName, _ := wsltestutils.RegisterDistro(t, ctx, true)
				d, err = db.GetDistroAndUpdateProperties(ctx, distroName, distro.Properties{})
				require.NoError(t, err, "Setup: GetDistroAndUpdateProperties should return no errors")
				defer d.Cleanup(ctx)
			}

			command := commandSetup(t, ctx, tc.command, d)
			if tc.command == cmdStop && !tc.mockErr {
				// We need to LockAwake, otherwise ReleaseAwake will error out
				require.NoError(t, d.LockAwake(), "Setup: could not lock distro awake")
			}

			client, err := landscape.NewClient(conf, db, landscape.WithHostname("HOSTNAME"))
			require.NoError(t, err, "Landscape NewClient should not return an error")

			err = client.Connect(ctx)
			require.NoError(t, err, "Setup: Connect should return no errors")
			defer client.Stop(ctx)

			require.Eventually(t, func() bool {
				return client.Connected() && conf.landscapeAgentUID != "" && service.IsConnected(conf.landscapeAgentUID)
			}, 10*time.Second, 100*time.Millisecond, "Landscape server and client never made a connection")

			enableMockErrors()

			err = service.SendCommand(ctx, conf.landscapeAgentUID, command)
			require.NoError(t, err, "SendCommand should return no error")

			// Allow some time for the message to be sent, received, and executed.
			time.Sleep(time.Second)
			if wsl.MockAvailable() && tc.command == cmdInstall || tc.command == cmdUninstall {
				// Appx state cannot be mocked
				return
			}

			if tc.command == cmdStop && tc.mockErr {
				// There is no way to assert on this function failing, as it is indistiguishable
				// from succeeding. I can fail two ways:
				//
				// - If Start was not called before. But the effect is the same as in success: the distro will be asleep.
				// - If the distro is no longer valid. Then the command takes no effect so we cannot assert on it.
				//
				// We still have the test case to exercise the code and ensure that it at least does not panic.
				return
			}

			// Disconnect to ensure the command has completed
			client.Stop(ctx)

			disableMockErrors()
			requireCommandResult(t, ctx, tc.command, d, conf, !tc.wantFailure)
		})
	}
}

const (
	testAppx       = "CanonicalGroupLimited.Ubuntu22.04LTS" // The name of the Appx
	testDistroAppx = "Ubuntu-22.04"                         // The name used in `wsl --install <DISTRO>`
)

//nolint:revive // testing.T goes before context
func commandSetup(t *testing.T, ctx context.Context, command command, distro *distro.Distro) *landscapeapi.Command {
	t.Helper()

	var r landscapeapi.Command

	switch command {
	case cmdAssignHost:
		r.Cmd = &landscapeapi.Command_AssignHost_{AssignHost: &landscapeapi.Command_AssignHost{Uid: "HostUID123"}}
	case cmdStart:
		r.Cmd = &landscapeapi.Command_Start_{Start: &landscapeapi.Command_Start{Id: distro.Name()}}
	case cmdStop:
		r.Cmd = &landscapeapi.Command_Stop_{Stop: &landscapeapi.Command_Stop{Id: distro.Name()}}
	case cmdInstall:
		r.Cmd = &landscapeapi.Command_Install_{Install: &landscapeapi.Command_Install{Id: testDistroAppx}}
		t.Cleanup(func() {
			d := wsl.NewDistro(ctx, testDistroAppx)
			_ = d.Uninstall(ctx)
		})
	case cmdUninstall:
		require.NoError(t, wsl.Install(ctx, testDistroAppx), "Setup: could not install Ubuntu-22.04")
		r.Cmd = &landscapeapi.Command_Uninstall_{Uninstall: &landscapeapi.Command_Uninstall{Id: testDistroAppx}}
		t.Cleanup(func() {
			d := wsl.NewDistro(ctx, testDistroAppx)
			_ = d.Uninstall(ctx)
		})
	case cmdSetDefault:
		otherDistro, _ := wsltestutils.RegisterDistro(t, ctx, false)
		d := wsl.NewDistro(ctx, otherDistro)
		err := d.SetAsDefault()
		require.NoError(t, err, "Setup: could not set another distro as default")
		r.Cmd = &landscapeapi.Command_SetDefault_{SetDefault: &landscapeapi.Command_SetDefault{Id: distro.Name()}}
	case cmdShutdownHost:
		d := wsl.NewDistro(ctx, distro.Name())
		err := d.Command(ctx, "exit 0").Run()
		require.NoError(t, err, "Setup: could not start distro")
		r.Cmd = &landscapeapi.Command_ShutdownHost_{ShutdownHost: &landscapeapi.Command_ShutdownHost{}}
	default:
		require.FailNowf(t, "Setup", "Unknown command type %d", command)
	}

	return &r
}

// requireCommandResult checks that a certain command has been executed in the machine
// by measuring its effect on the targeted distro. Set wantSuccess to true if you want
// to assert that the command completed successfully, and set it to false to assert it
// did not complete.
//
//nolint:revive // testing.T goes before context
func requireCommandResult(t *testing.T, ctx context.Context, command command, distro *distro.Distro, conf *mockConfig, wantSuccess bool) {
	t.Helper()

	// Ensuring a connection was made
	require.NotEmpty(t, conf.landscapeAgentUID, "Landscape client should have had a UID assigned by the server")

	switch command {
	case cmdAssignHost:
		if wantSuccess {
			require.Equal(t, "HostUID123", conf.landscapeAgentUID, "Landscape client should have overridden the initial UID sent by the server")
		} else {
			require.Fail(t, "Test not implemented")
		}
	case cmdStart:
		ok, state := checkEventuallyState(t, distro, wsl.Running, 10*time.Second, time.Second)
		if wantSuccess {
			require.True(t, ok, "Distro never reached %q state. Last state: %q", wsl.Running, state)
		} else {
			require.False(t, ok, "Distro unexpectedly reached state %q", wsl.Running)
		}
	case cmdStop:
		// We wait a bit longer than WSL sleep time, because we must account for the Landscape server-client
		// interaction completing asyncronously with the test.
		const waitFor = 15 * time.Second
		ok, state := checkEventuallyState(t, distro, wsl.Stopped, waitFor, time.Second)
		if wantSuccess {
			require.True(t, ok, "Distro never reached %q state. Last state: %q", wsl.Running, state)
		} else {
			require.False(t, ok, "Distro unexpectedly reached state %q", wsl.Stopped)
		}
	case cmdInstall:
		inst := isAppxInstalled(t, testAppx)

		d := wsl.NewDistro(ctx, distro.Name())
		reg, err := d.IsRegistered()
		require.NoError(t, err, "IsRegistered should return no error")

		if wantSuccess {
			require.True(t, inst, "Appx should have been installed, but it wasn't")
			require.True(t, reg, "Distro should have been registered")

			conf, err := d.GetConfiguration()
			require.NoError(t, err, "GetConfiguration should return no error")
			require.NotEqual(t, uint32(0), conf.DefaultUID, "Default user should have been changed from root")
		} else {
			require.False(t, inst, "Appx should not have been installed")
			require.True(t, reg, "Distro should not have been registered")
		}
	case cmdUninstall:
		inst := isAppxInstalled(t, testAppx)
		if wantSuccess {
			require.False(t, inst, "Appx should no longer be installed, but it is")
		} else {
			require.True(t, inst, "Appx should still be installed, but it isn't")
		}
	case cmdSetDefault:
		def, err := wsl.DefaultDistro(ctx)
		require.NoError(t, err, "could not call DefaultDistro")
		if wantSuccess {
			require.Equal(t, distro.Name(), def.Name(), "Test distro should be the default one")
		} else {
			require.NotEqual(t, distro.Name(), def.Name(), "Test distro should not be the default one")
		}
	case cmdShutdownHost:
		gotState, err := distro.State()
		require.NoError(t, err, "Could not read distro state")
		if wantSuccess {
			require.Equal(t, wsl.Stopped, gotState, "Unexpected disto state. Want: %q. Got: %q", wsl.Stopped, gotState)
		} else {
			require.Equal(t, wsl.Running, gotState, "Unexpected disto state. Want: %q. Got: %q", wsl.Running, gotState)
		}
	default:
		require.FailNowf(t, "Setup", "Unknown command type %d", command)
	}
}

func checkEventuallyState(t *testing.T, d *distro.Distro, wantState wsl.State, waitFor, tick time.Duration) (ok bool, lastState wsl.State) {
	t.Helper()

	timer := time.NewTimer(waitFor)
	defer timer.Stop()

	ticker := time.NewTicker(tick)
	defer ticker.Stop()

	for {
		select {
		case <-timer.C:
			return false, lastState
		case <-ticker.C:
			var err error
			lastState, err = d.State()
			require.NoError(t, err, "disto State should return no error")
			if lastState == wantState {
				return true, lastState
			}
		}
	}
}

func isAppxInstalled(t *testing.T, appxPackage string) bool {
	t.Helper()
	require.False(t, wsl.MockAvailable(), "This assertion is only valid without the WSL mock")

	cmd := fmt.Sprintf("(Get-AppxPackage -Name %q).Status", appxPackage)
	//nolint:gosec // Command with variable is acceptable in test code
	out, err := exec.Command("powershell.exe", "-NoProfile", "-NoLogo", "-NonInteractive", "-Command", cmd).Output()
	require.NoError(t, err, "Get-AppxPackage should return no error. Stdout: %s", string(out))

	return strings.Contains(string(out), "Ok")
}

//nolint:revive // Context goes after testing.T
func setUpLandscapeMock(t *testing.T, ctx context.Context, addr string, requireCertificate bool) (lis net.Listener, server *grpc.Server, service *landscapemockservice.Service) {
	t.Helper()

	var cfg net.ListenConfig
	lis, err := cfg.Listen(ctx, "tcp", addr)
	require.NoError(t, err, "Setup: can't listen")

	var opts []grpc.ServerOption
	if requireCertificate {
		certPath := filepath.Join(golden.TestFamilyPath(t), "certificates/cert.pem")
		keyPath := filepath.Join(golden.TestFamilyPath(t), "certificates/key.pem")

		serverCert, err := tls.LoadX509KeyPair(certPath, keyPath)
		require.NoError(t, err, "Setup: could not load Landscape mock server credentials")

		config := &tls.Config{
			Certificates: []tls.Certificate{serverCert},
			ClientAuth:   tls.NoClientCert,
			MinVersion:   tls.VersionTLS12,
		}

		opts = append(opts, grpc.Creds(credentials.NewTLS(config)))
	}

	server = grpc.NewServer(opts...)
	service = landscapemockservice.New()
	landscapeapi.RegisterLandscapeHostAgentServer(server, service)

	return lis, server, service
}

type mockConfig struct {
	proToken              string
	landscapeClientConfig string
	landscapeAgentUID     string

	proTokenErr        bool
	landscapeConfigErr bool
	landscapeUIDErr    bool
	setLandscapeUIDErr bool

	mu sync.Mutex
}

func (m *mockConfig) LandscapeClientConfig(ctx context.Context) (string, config.Source, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.landscapeConfigErr {
		return "", config.SourceNone, errors.New("Mock error")
	}
	return m.landscapeClientConfig, config.SourceUser, nil
}

func (m *mockConfig) ProvisioningTasks(ctx context.Context, distroName string) ([]task.Task, error) {
	return nil, nil
}

func (m *mockConfig) Subscription(ctx context.Context) (string, config.Source, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.proTokenErr {
		return "", config.SourceNone, errors.New("Mock error")
	}
	return m.proToken, config.SourceUser, nil
}

func (m *mockConfig) LandscapeAgentUID(ctx context.Context) (string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.landscapeUIDErr {
		return "", errors.New("Mock error")
	}
	return m.landscapeAgentUID, nil
}

func (m *mockConfig) SetLandscapeAgentUID(ctx context.Context, uid string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.setLandscapeUIDErr {
		return errors.New("Mock error")
	}

	m.landscapeAgentUID = uid
	return nil
}
