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
	"github.com/endorama/devid/cmd/utils"
	"github.com/endorama/devid/internal/backup"
)

func Backup() *cobra.Command {
	// backupCmd represents the backup command.
	var backupCmd = &cobra.Command{ //nolint:gochecknoglobals // required by cobra
		Use:   "backup",
		Short: "backup a persona",
		Long: `Create encrypted backup of a persona.

The backup is compressed (.tar.gz) and encrypted using age (filippo.io/age).
Encryption requires a passphrase that is automatically generated using a safe 
RNG function and printed after backup creation.

This command loads the current persona from DEVID_ACTIVE_PERSONA environment variable, and this 
value takes precedence over the --persona flag.
`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			p, err := utils.LoadPersona(cmd)
			if err != nil {
				ui.Fatal(fmt.Errorf("cannot instantiate persona: %w", err), noPersonaLoadedExitCode)
			}
			passphrase := utils.GeneratePassphrase()

			ui.Outputf(fmt.Sprintf("creating backup for persona: %s\n", p.Name()))

			out, err := os.Create(fmt.Sprintf("%s.tar.gz.age", p.Name()))
			if err != nil {
				ui.Fatal(fmt.Errorf("cannot create file: %w", err), genericExitCode)
			}
			defer out.Close()

			b, err := backup.NewTask(p.Name(), p.Location(), out)
			if err != nil {
				ui.Fatal(fmt.Errorf("cannot create backup task: %w", err), genericExitCode)
			}
			err = backup.Perform(b, passphrase)
			if err != nil {
				ui.Fatal(err, genericExitCode)
			}

			ui.Infof("Encryption passphrase is: %s", passphrase)
		},
	}
	backupCmd.Flags().String("persona", "", "The persona to backup")

	return backupCmd

}
