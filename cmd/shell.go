package cmd

import (
	"errors"
	"fmt"
	"os"
	"path"
	"syscall"

	"github.com/endorama/devid/cmd/ui"
	"github.com/endorama/devid/cmd/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var shellCmd = &cobra.Command{ //nolint:gochecknoglobals // required by cobra
	Use:   "shell",
	Short: "load a shell preconfigured with persona environment",
	Long: `Execute the load.sh file of the specified persona, loading the environment,

This command loads the current persona from DEVID_ACTIVE_PERSONA environment variable, and this value takes precedence over the --persona flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		p, err := utils.LoadPersona(cmd)
		if err != nil {
			ui.Fatal(fmt.Errorf("cannot instantiate persona: %w", err), noPersonaLoadedExitCode)
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
	shellCmd.Flags().String("persona", "", "The persona's shell to load")
	rootCmd.AddCommand(shellCmd)
}
