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

	"github.com/spf13/cobra"

	"github.com/endorama/devid/internal/plugin/manager"
)

// listCmd represents the list command.
var pluginsCmd = &cobra.Command{ //nolint:gochecknoglobals // required by cobra
	Use:   "plugins",
	Short: "List all bundled plugins",
	Long: `devid plugins

List all bundled plugins, Core and Optional.
Plugins are executed in the order they are displayed on this list.
Core plugins are always enabled.
Optional plugins must be enabled in Persona config file.
`,
	Run: func(cmd *cobra.Command, args []string) {
		ui.Output("Core plugins:")
		for _, k := range manager.Core {
			ui.Output(fmt.Sprintf("- %s", k.Name()))
		}
		ui.Output("Optional plugins:")
		for _, k := range manager.Optional {
			ui.Output(fmt.Sprintf("- %s", k.Name()))
		}
	},
}

func init() { //nolint:gochecknoinits // required by cobra
	rootCmd.AddCommand(pluginsCmd)
}
