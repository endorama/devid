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
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/endorama/devid/internal/persona"
	"github.com/endorama/devid/internal/plugin"
	"github.com/endorama/devid/internal/plugin/manager"
	"github.com/endorama/devid/internal/utils"
)

// rehashCmd represents the rehash command.
var rehashCmd = &cobra.Command{ //nolint:gochecknoglobals // required by cobra
	Use:   "rehash",
	Short: "Rebuild profiles loader and shims",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if os.Getenv("DEVID_ACTIVE_PROFILE") != "" {
			// NOTE: rehashing when a profile is active is dangerous, as the environment
			// has been changed with customizations and there is no guarantee about
			// what those changes have affected.
			// This may be especially problematic for executable path detection in
			// plugin.
			// As such we prevent rehashing while there is an active profile.
			ui.Error("Trying to rehash with an active profile. This may go very wrong.")
			os.Exit(1)
		}

		currentPersona, err := cmd.Flags().GetString("persona")
		if err != nil {
			ui.Error(fmt.Errorf("cannot access flag currentPersona: %w", err).Error())
		}
		if currentPersona != "" {
			p, err := persona.New(currentPersona)
			if err != nil {
				ui.Error(fmt.Errorf("cannot instantiate persona: %w", err).Error())
				os.Exit(1)
			}
			if !p.Exists() {
				ui.Error("persona does not exists")
			}

			config, err := plugin.LoadConfigFromFile(p.File())
			if err != nil {
				ui.Error(fmt.Errorf("cannot load configuration from file (%s): %w", p.File(), err).Error())
				os.Exit(1)
			}

			p.Config = config

			errs, err := manager.LoadCorePlugins(p.Config)
			if err != nil {
				ui.Error(err.Error())

				for _, e := range errs {
					ui.Error(e.Error())
				}

				os.Exit(pluginManagerCoreLoadingErrorExitCode)
			}

			log.Printf("%+v\n", p)

			content, err := manager.ShellLoader(p)
			if err != nil {
				ui.Error(err.Error())
				os.Exit(1)
			}

			log.Printf("%+v\n", content)

			shellLoaderFilePath := path.Join(p.Location(), viper.GetString("shell_loader_filename"))
			err = utils.PersistFile(shellLoaderFilePath, content)
			if err != nil {
				ui.Error(fmt.Errorf("cannot save shell loader: %w", err).Error())
				os.Exit(1)
			}
		} else {
			ui.Error("Not yet implemented")
			os.Exit(1)
		}
	},
}

func init() { //nolint:gochecknoinits // required by cobra
	rootCmd.AddCommand(rehashCmd)
	rehashCmd.Flags().String("persona", "", "The persona to backup")
}
