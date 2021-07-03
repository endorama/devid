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
	"os"

	"github.com/spf13/cobra"

	"github.com/endorama/devid/internal/persona"
	"github.com/endorama/devid/internal/plugin/manager"
	"github.com/endorama/devid/internal/utils"
)

// newCmd represents the new command.
var newCmd = &cobra.Command{ //nolint:gochecknoglobals // required by cobra
	Use:   "new",
	Short: "Create a new (empty) persona",
	Long: fmt.Sprintf(`devid new <persona name>

Create a new persona configuration file, opens it within EDITOR.

Allowed EDITOR values: %s
`, utils.AllowedEditors),
	Run: func(cmd *cobra.Command, args []string) {
		runCommand(args)
	},
}

func init() { //nolint:gochecknoinits // required by cobra
	rootCmd.AddCommand(newCmd)

	// add --overwrite to overwrite already existing profile
}

func runCommand(args []string) {
	if len(args) != 1 {
		ui.Error("Argument NAME required")
		os.Exit(genericExitCode)
	}

	name := args[0]

	p, _ := persona.New(name)

	err := persona.Create(p)
	if err != nil {
		ui.Error(err.Error())
		os.Exit(genericExitCode)
	}

	errs, err := manager.LoadCorePlugins(p.Config)
	if err != nil {
		ui.Error(err.Error())

		for _, e := range errs {
			ui.Error(e.Error())
		}

		os.Exit(pluginManagerCoreLoadingErrorExitCode)
	}

	errs, err = manager.SetupPlugins(p)
	if err != nil {
		ui.Error(err.Error())

		for _, e := range errs {
			ui.Error(e.Error())
		}

		os.Exit(pluginManagerSetupErrorExitCode)
	}

	err = utils.OpenWithEditor(p.File())
	if err != nil {
		// deleting the created persona so new does not error the second time is run
		_ = persona.Delete(p)

		ui.Error(err.Error())
		os.Exit(genericExitCode)
	}
}
