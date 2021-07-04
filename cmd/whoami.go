/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/endorama/devid/internal/persona"
	"github.com/endorama/devid/internal/plugin"
	"github.com/endorama/devid/internal/plugin/manager"
	"github.com/endorama/devid/plugins/identity"
)

// whoamiCmd represents the whoami command.
var whoamiCmd = &cobra.Command{ //nolint:gochecknoglobals // required by cobra
	Use:   "whoami",
	Short: "Print current loaded persona",
	Long: `Print current loaded persona.

If no persona is loaded print nothing and exit with code 128.
`,
	Run: func(cmd *cobra.Command, args []string) {
		currentPersona := viper.GetString("active_persona")
		log.Println(currentPersona)

		// there is no loaded profile
		if currentPersona == "" {
			os.Exit(noPersonalLoadedExitCode)
		}

		extended, err := cmd.Flags().GetBool("extended")
		if err != nil {
			ui.Error(fmt.Errorf("cannot access flag extended: %w", err).Error())
		}
		if extended {
			p, _ := persona.New(currentPersona)

			config, err := plugin.LoadConfigFromFile(p.File())
			if err != nil {
				ui.Error(fmt.Errorf("cannot load configuration from file (%s): %w", p.File(), err).Error())
				os.Exit(1)
			}

			p.Config = config
			manager.LoadCorePlugins(p.Config)

			plg, found := manager.GetPlugin("identity")
			if identityPlugin, ok := plg.Instance.(*identity.Plugin); found && ok {
				currentPersona = fmt.Sprintf("%s, %s", currentPersona, identityPlugin.Whoami())
			}
		}

		ui.Output(currentPersona)
		os.Exit(0)
	},
}

func init() { //nolint:gochecknoinits // required by cobra
	rootCmd.AddCommand(whoamiCmd)

	whoamiCmd.Flags().BoolP("extended", "e", false, "Print extended identity information")
}
