// Package restserver provides building blocks to implement a mocked version of out-of-process components Ubuntu Pro For Windows depend on that talk REST,
// such as MS Store API and the Contracts Server backend
// DO NOT USE IN PRODUCTION
package restserver

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// Server is the minimal interface mock REST servers must provide to the Application.
type Server interface {
	Stop() error
	Serve(ctx context.Context) (string, error)
}

// Settings is the minimal interface a settings backend must provide to the Application.
type Settings interface {
	SetAddress(address string)
	GetAddress() string
}

// Application encapsulates creating and managing the CLI and lifecycle.
type Application struct {
	// Name of the application as shown in the help messages.
	Name string
	// Description fo the application as shown in the long help messages.
	Description string
	// The default settings that the application will pass to the server instance.
	DefaultSettings Settings
	// A function capable of translating from the Settings interface into the concrete implementation a particular server will need to run.
	ServerFactory func(Settings) Server
}

// Execute runs the server CLI.
func (a *Application) Execute() int {
	rootCmd := a.rootCmd()
	rootCmd.AddCommand(a.showDefaultsCmd())

	rootCmd.PersistentFlags().CountP("verbosity", "v", "WARNING (-v) INFO (-vv), DEBUG (-vvv)")
	rootCmd.PersistentFlags().StringP("output", "o", "", "File where relevant non-log output will be written to")
	rootCmd.Flags().StringP("address", "a", "", "Overrides the address where the server will be hosted")

	if err := rootCmd.Execute(); err != nil {
		slog.Error(fmt.Sprintf("Error executing: %v", err))
		return 1
	}
	return 0
}

// setVerboseMode changes the verbosity of the logs.
func setVerboseMode(n int) {
	var level slog.Level
	switch n {
	case 0:
		level = slog.LevelError
	case 1:
		level = slog.LevelWarn
	case 2:
		level = slog.LevelInfo
	default:
		level = slog.LevelDebug
	}

	h := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	slog.SetDefault(slog.New(h))
}

func execName() string {
	exe, err := os.Executable()
	if err != nil {
		slog.Error(fmt.Sprintf("Could not get executable name: %v", err))
		os.Exit(1)
	}

	return filepath.Base(exe)
}

func (a *Application) showDefaultsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "show-defaults",
		Short: fmt.Sprintf("See the default values for the %s server", a.Name),
		Long:  fmt.Sprintf("See the default values for the %s server. These are the settings that 'serve' will use unless overridden.", a.Description),
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			out, err := yaml.Marshal(a.DefaultSettings)
			if err != nil {
				slog.Error(fmt.Sprintf("Could not marshal default settings: %v", err))
				os.Exit(1)
			}

			if outfile := cmd.Flag("output").Value.String(); outfile != "" {
				if err := os.WriteFile(outfile, out, 0600); err != nil {
					slog.Error(fmt.Sprintf("Could not write to output file: %v", err))
					os.Exit(1)
				}
				return
			}

			fmt.Println(string(out))
		},
	}
}

func (a *Application) rootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   fmt.Sprintf("%s [settings_file]", execName()),
		Short: fmt.Sprintf("A mock %s server for Ubuntu Pro For Windows testing", a.Name),
		Long: fmt.Sprintf(`A mock of the %s for Ubuntu Pro For Windows testing.
Serve the store server with the optional settings file.
Default settings will be used if none are provided.
The outfile, if provided, will contain the address.`, a.Description),
		Args: cobra.RangeArgs(0, 1),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Force a visit of the local flags so persistent flags for all parents are merged.
			cmd.LocalFlags()

			// command parsing has been successful. Returns to not print usage anymore.
			cmd.SilenceUsage = true

			v := cmd.Flag("verbosity").Value.String()
			n, err := strconv.Atoi(v)
			if err != nil {
				return fmt.Errorf("could not parse verbosity: %v", err)
			}

			setVerboseMode(n)
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			settings := a.DefaultSettings

			if len(args) > 0 {
				out, err := os.ReadFile(args[0])
				if err != nil {
					slog.Error(fmt.Sprintf("Could not read input file %q: %v", args[0], err))
					os.Exit(1)
				}

				if err := yaml.Unmarshal(out, &settings); err != nil {
					slog.Error(fmt.Sprintf("Could not unmarshal settings: %v", err))
					os.Exit(1)
				}
			}

			if addr := cmd.Flag("address").Value.String(); addr != "" {
				settings.SetAddress(addr)
			}

			sv := a.ServerFactory(settings)
			addr, err := sv.Serve(ctx)
			if err != nil {
				slog.Error(fmt.Sprintf("Could not serve: %v", err))
				os.Exit(1)
			}

			defer func() {
				if err := sv.Stop(); err != nil {
					slog.Error(fmt.Sprintf("stopped serving: %v", err))
				}
				slog.Info("stopped serving")
			}()

			if outfile := cmd.Flag("output").Value.String(); outfile != "" {
				if err := os.WriteFile(outfile, []byte(addr), 0600); err != nil {
					slog.Error(fmt.Sprintf("Could not write output file: %v", err))
					os.Exit(1)
				}
			}

			slog.Info(fmt.Sprintf("Serving on address %s", addr))

			// Wait loop
			for scanned := ""; scanned != "exit"; fmt.Scanf("%s\n", &scanned) {
				fmt.Println("Write 'exit' to stop serving")
			}
		},
	}
}
