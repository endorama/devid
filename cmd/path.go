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

	"github.com/endorama/devid/cmd/ui"
	"github.com/endorama/devid/cmd/utils"
)

// pathCmd represent the path sub-command.
func Path() *cobra.Command {
	pathCmd := &cobra.Command{
		Use:   "path",
		Short: "show persona location on disk",
		Long:  `Prints the location on disk of persona configuration folder.`,
		Run: func(cmd *cobra.Command, _ []string) {
			p, err := utils.LoadPersona(cmd)
			if err != nil {
				ui.Fatal(fmt.Errorf("cannot instantiate persona: %w", err), noPersonaLoadedExitCode)
			}

			ui.Outputf(p.Location())
		},
	}

	pathCmd.Flags().String("persona", "", "The persona to delete")

	return pathCmd
}
