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
	cmdutils "github.com/endorama/devid/cmd/utils"
	"github.com/endorama/devid/internal/utils"
)

func Edit() *cobra.Command {
	editCmd := &cobra.Command{ //nolint:gochecknoglobals // required by cobra
		Use:   "edit",
		Short: "edit a persona definition file in your $EDITOR",
		Long: fmt.Sprintf(`Open within EDITOR the specified persona configuration file.

For security EDITOR variable content is matched against a list of valid editor executable paths.
NOTE however that if some of this commands are not available on your system is still possible to 
trigger an unknown command execution trough this command.

Allowed EDITOR values: %s

This command loads the current persona from DEVID_ACTIVE_PERSONA environment variable, and this 
value takes precedence over the --persona flag.
`, utils.AllowedEditors),
		Run: func(cmd *cobra.Command, _ []string) {
			p, err := cmdutils.LoadPersona(cmd)
			if err != nil {
				ui.Fatal(fmt.Errorf("cannot instantiate persona: %w", err), genericExitCode)
			}

			err = utils.OpenWithEditor(p.File())
			if err != nil {
				ui.Fatal(err, genericExitCode)
			}
		},
	}

	editCmd.Flags().String("persona", "", "The persona to backup")

	return editCmd
}
