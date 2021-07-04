package cmd

import (
	"fmt"
	"os"
	"path"
	"syscall"

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
			ui.Error(fmt.Errorf("cannot access flag currentPersona: %w", err).Error())
		}
		if currentPersona == "" {
			ui.Error("--persona requires a value")
			os.Exit(1)
		}

		p, err := persona.New(currentPersona)
		if err != nil {
			ui.Error(fmt.Errorf("cannot instantiate persona: %w", err).Error())
			os.Exit(1)
		}
		if !p.Exists() {
			ui.Error("persona does not exists")
		}

		shellLoaderFilePath := path.Join(p.Location(), viper.GetString("shell_loader_filename"))

		err = syscall.Exec("/usr/bin/env", []string{"bash", shellLoaderFilePath}, os.Environ())
		if err != nil {
			ui.Error(fmt.Errorf("cannot exec load file: %w", err).Error())
			os.Exit(1)
		}
	},
}

func init() { //nolint:gochecknoinits // required by cobra
	rootCmd.AddCommand(shellCmd)
	shellCmd.Flags().String("persona", "", "The persona to backup")
}
