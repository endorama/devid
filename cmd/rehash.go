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
	"errors"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/endorama/devid/cmd/ui"
	cmdutils "github.com/endorama/devid/cmd/utils"
	"github.com/endorama/devid/internal/plugin/manager"
	"github.com/endorama/devid/internal/utils"
)

// rehashCmd represents the rehash command.
var rehashCmd = &cobra.Command{ //nolint:gochecknoglobals // required by cobra
	Use:   "rehash",
	Short: "Rebuild profiles loader and shims",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if viper.GetString("active_persona") != "" {
			// NOTE: rehashing when a profile is active is dangerous, as the environment
			// has been changed with customizations and there is no guarantee about
			// what those changes have affected.
			// This may be especially problematic for executable path detection in
			// plugin.
			// As such we prevent rehashing while there is an active profile.
			err := errors.New("Trying to rehash with an active profile. This may go very wrong.")
			ui.Fatal(err, genericExitCode)
		}

		p, err := cmdutils.LoadPersona(cmd)
		if err != nil {
			ui.Fatal(fmt.Errorf("cannot instantiate persona: %w", err), noPersonaLoadedExitCode)
		}
		// errs, err := manager.LoadPlugins(p.Config)
		// if err != nil {
		//   ui.Error(err.Error())
		//
		//   for _, e := range errs {
		//     ui.Error(e.Error())
		//   }
		//
		//   os.Exit(pluginManagerLoadingErrorExitCode)
		// }

		errs, err := manager.LoadCorePlugins(p.Config)
		if err != nil {
			ui.Error(err)

			for _, e := range errs {
				ui.Error(e)
			}

			os.Exit(pluginManagerCoreLoadingErrorExitCode)
		}

		errs, err = manager.LoadOptionalPlugins(p.Config)
		if err != nil {
			ui.Error(err)

			for _, e := range errs {
				ui.Error(e)
			}

			os.Exit(pluginManagerOptionalLoadingErrorExitCode)
		}

		log.Printf("persona: %+v\n", p)

		errs, err = manager.Generate(p)
		if err != nil {
			ui.Error(err)

			for _, e := range errs {
				ui.Error(e)
			}

			os.Exit(pluginGenerationExitCode)
		}

		content, err := manager.ShellLoader(p)
		if err != nil {
			ui.Error(err)
			os.Exit(1)
		}

		log.Printf("%+v\n", content)

		shellLoaderFilePath := path.Join(p.Location(), viper.GetString("shell_loader_filename"))
		err = utils.PersistFile(shellLoaderFilePath, content)
		if err != nil {
			ui.Fatal(fmt.Errorf("cannot save shell loader: %w", err), genericExitCode)
		}
		os.Chmod(shellLoaderFilePath, 0700)
	},
}

func init() { //nolint:gochecknoinits // required by cobra
	rootCmd.AddCommand(rehashCmd)
	rehashCmd.Flags().String("persona", "", "The persona for which to rebuild the shell configuration")
}
