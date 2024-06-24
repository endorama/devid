package cmd

import (
	"fmt"
	"os"
	"path"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/endorama/devid/cmd/ui"
	"github.com/endorama/devid/cmd/utils"
)

func Shell() *cobra.Command {
	shellCmd := &cobra.Command{ //nolint:gochecknoglobals // required by cobra
		Use:   "shell",
		Short: "load a shell preconfigured with persona environment",
		Long: `Execute the load.sh file of the specified persona, loading the environment,

This command loads the current persona from DEVID_ACTIVE_PERSONA environment variable, and this 
value takes precedence over the --persona flag.`,
		Run: func(cmd *cobra.Command, _ []string) {
			p, err := utils.LoadPersona(cmd)
			if err != nil {
				ui.Fatal(fmt.Errorf("cannot instantiate persona: %w", err), noPersonaLoadedExitCode)
			}
			if !p.Exists() {
				ui.Fatal(errPersonaDontExists, genericExitCode)
			}

			shellLoaderFilePath := path.Join(p.Location(), viper.GetString("shell_loader_filename"))

			//#nosec G204 -- shellLoaderFilePath is not user controlled (check shell_loader_filename settings)
			err = syscall.Exec("/usr/bin/env", []string{"bash", shellLoaderFilePath}, os.Environ())
			if err != nil {
				ui.Fatal(fmt.Errorf("cannot exec load file: %w", err), genericExitCode)
			}
		},
	}

	shellCmd.Flags().String("persona", "", "The persona's shell to load")

	return shellCmd
}
