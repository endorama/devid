package cmd

import (
	"errors"
	"fmt"
	"os"
	"path"
	"syscall"

	"github.com/endorama/devid/cmd/ui"
	"github.com/endorama/devid/internal/persona"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var shellCmd = &cobra.Command{ //nolint:gochecknoglobals // required by cobra
	Use:   "shell",
	Short: "Load a shell preconfigured with persona environment",
	Long: `devid shell --persona=<name>

exec the load.sh file of the specified persona, loading the environment`,
	Run: func(cmd *cobra.Command, args []string) {
		currentPersona, err := cmd.Flags().GetString("persona")
		if err != nil {
			ui.Fatal(fmt.Errorf("cannot access flag currentPersona: %w", err), genericExitCode)

		}
		if currentPersona == "" {
			err := errors.New("--persona requires a value")
			ui.Fatal(err, genericExitCode)
		}

		p, err := persona.New(currentPersona)
		if err != nil {
			ui.Fatal(fmt.Errorf("cannot instantiate persona: %w", err), genericExitCode)
		}
		if !p.Exists() {
			err := errors.New("persona does not exists")
			ui.Fatal(err, genericExitCode)
		}

		shellLoaderFilePath := path.Join(p.Location(), viper.GetString("shell_loader_filename"))

		err = syscall.Exec("/usr/bin/env", []string{"bash", shellLoaderFilePath}, os.Environ())
		if err != nil {
			ui.Fatal(fmt.Errorf("cannot exec load file: %w", err), genericExitCode)
		}
	},
}

func init() { //nolint:gochecknoinits // required by cobra
	rootCmd.AddCommand(shellCmd)
	shellCmd.Flags().String("persona", "", "The persona to backup")
}
