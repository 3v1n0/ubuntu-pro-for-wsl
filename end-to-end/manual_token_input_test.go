package endtoend_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	wsl "github.com/ubuntu/gowsl"
)

func TestManualTokenInput(t *testing.T) {
	type whenToken int
	const (
		never whenToken = iota
		beforeDistroRegistration
		afterDistroRegistration
	)

	// Let's be lazy and don't fall into the risk of changing the function name without updating the places where its name is used.
	currentFuncName := t.Name()

	testCases := map[string]struct {
		whenToken        whenToken
		overrideTokenEnv string

		wantAttached bool
	}{
		"Success when applying pro token before registration": {whenToken: beforeDistroRegistration, wantAttached: true},
		"Success when applying pro token after registration":  {whenToken: afterDistroRegistration, wantAttached: true},

		"Error with invalid token": {whenToken: afterDistroRegistration, overrideTokenEnv: fmt.Sprintf("%s=%s", proTokenEnv, "CJd8MMN8wXSWsv7wJT8c8dDK")},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()

			testSetup(t)

			// Either runs the ubuntupro app before...
			if tc.whenToken == beforeDistroRegistration {
				cleanup := startAgent(t, ctx, currentFuncName, tc.overrideTokenEnv)
				defer cleanup()
			}

			// Distro setup
			name := registerFromTestImage(t, ctx)
			d := wsl.NewDistro(ctx, name)

			defer logWslProServiceJournal(t, ctx, d)

			out, err := d.Command(ctx, "exit 0").CombinedOutput()
			require.NoErrorf(t, err, "Setup: could not wake distro up: %v. %s", err, out)

			// ... or after registration, but never both.
			if tc.whenToken == afterDistroRegistration {
				// the exact time between registering the distro and running the app cannot be determined, so there is a chance
				// of this starting the app before registration completes.
				time.Sleep(5 * time.Second)
				cleanup := startAgent(t, ctx, currentFuncName, tc.overrideTokenEnv)
				defer cleanup()
			}

			time.Sleep(5 * time.Second)
			out, err = d.Command(ctx, "exit 0").CombinedOutput()
			require.NoErrorf(t, err, "Setup: could not wake distro up: %v. %s", err, out)

			maxTimeout := 30 * time.Second
			if !tc.wantAttached {
				time.Sleep(maxTimeout)
				attached, err := distroIsProAttached(t, d)
				require.NoError(t, err, "could not determine if distro is attached")
				require.False(t, attached, "distro should not have been Pro attached")
				return
			}

			require.Eventually(t, func() bool {
				attached, err := distroIsProAttached(t, d)
				if err != nil {
					t.Logf("could not determine if distro is attached: %v", err)
				}
				return attached
			}, maxTimeout, time.Second, "distro should have been Pro attached")
		})
	}
}
