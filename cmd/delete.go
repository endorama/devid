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

	"github.com/lu4p/shred"
	"github.com/spf13/cobra"

	"github.com/endorama/devid/cmd/ui"
	"github.com/endorama/devid/cmd/utils"
)

func Delete() *cobra.Command {
	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "delete a persona",
		Long: `Delete a persona using a secure deletion function.

All content from the persona's folder will be destroyed through this command.

Deletion is performed using a Golang implementation of Linux shred command, configured to perform 3
iteration, overriding with zero and final file removal.

Note that shred relies on a very important assumption: that the file system overwrites data in 
place. This may not be the case for your file system, please refer to shred documentation for 
further details.

This command loads the current persona from DEVID_ACTIVE_PERSONA environment variable, and this 
value takes precedence over the --persona flag.
`,
		Run: func(cmd *cobra.Command, args []string) {
			p, err := utils.LoadPersona(cmd)
			if err != nil {
				ui.Fatal(fmt.Errorf("cannot instantiate persona: %w", err), noPersonaLoadedExitCode)
			}

			// TODO: ask for confirmation (and add it to command docs)

			shredTimes := 3
			shredconf := shred.Conf{Times: shredTimes, Zeros: true, Remove: true}
			err = shredconf.Dir(p.Location())
			if err != nil {
				ui.Fatal(err, genericExitCode)
			}
		},
	}

	deleteCmd.Flags().String("persona", "", "The persona to delete")

	return deleteCmd
}
