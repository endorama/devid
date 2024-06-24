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

	"github.com/endorama/devid/cmd/ui"
	"github.com/endorama/devid/internal/persona"
	"github.com/endorama/devid/internal/plugin/manager"
	"github.com/endorama/devid/internal/utils"
)

// newCmd represents the new command.
func New() *cobra.Command {
	newCmd := &cobra.Command{ //nolint:gochecknoglobals // required by cobra
		Use:   "new [name]",
		Short: "create a persona",
		Long: fmt.Sprintf(`Create a new, empty, persona configuration file, opens it within EDITOR.

For security EDITOR variable content is matched against a list of valid editor executable paths.
NOTE however that if some of this commands are not available on your system is still possible to 
trigger an unknown command execution trough this command.

Allowed EDITOR values: %s
`, utils.AllowedEditors),
		Args: cobra.MinimumNArgs(1),
		Run: func(_ *cobra.Command, args []string) {
			runCommand(args)
		},
	}

	// add --overwrite to overwrite already existing profile
	return newCmd
}

func runCommand(args []string) {
	name := args[0]

	p, _ := persona.New(name)

	err := persona.Create(p)
	if err != nil {
		ui.Fatal(err, genericExitCode)
	}

	errs, err := manager.LoadCorePlugins(p.Config)
	if err != nil {
		ui.Error(err)

		for _, e := range errs {
			ui.Error(e)
		}

		// deleting the created persona so new does not error the second time is run
		// _ = persona.Delete(p)
		os.Exit(pluginManagerCoreLoadingErrorExitCode)
	}

	errs, err = manager.LoadOptionalPlugins(p.Config)
	if err != nil {
		ui.Error(err)

		for _, e := range errs {
			ui.Error(e)
		}

		// deleting the created persona so new does not error the second time is run
		_ = persona.Delete(p)

		os.Exit(pluginManagerOptionalLoadingErrorExitCode)
	}

	errs, err = manager.SetupPlugins(p)
	if err != nil {
		ui.Error(err)

		for _, e := range errs {
			ui.Error(e)
		}

		// deleting the created persona so new does not error the second time is run
		_ = persona.Delete(p)

		os.Exit(pluginManagerSetupErrorExitCode)
	}

	err = utils.OpenWithEditor(p.File())
	if err != nil {
		// deleting the created persona so new does not error the second time is run
		_ = persona.Delete(p)

		ui.Fatal(err, genericExitCode)
	}
}
