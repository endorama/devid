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

	petname "github.com/dustinkirkland/golang-petname"
	"github.com/endorama/devid/internal/backup"
	"github.com/endorama/devid/internal/persona"
	"github.com/spf13/cobra"
)

var (
	currentPersona string
)

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Create encrypted backup of personas",
	Long: `Create encrypted backup of personas.

The backup is compressed (.tar.gz) and encrypted using age (filippo.io/age).
Encryption requires a passphrase that is automatically generated using a safe RNG function and printed after backup creation.
`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		ui.Output("backup called")
		if currentPersona != "" {
			petname.NonDeterministicMode()
			passphraseLength := 6
			passphrase := petname.Generate(passphraseLength, "-")

			p, _ := persona.New(currentPersona)
			if !p.Exists() {
				log.Fatalf("persona does not exists")
			}
			ui.Output(fmt.Sprintf("Creating backup for persona: %s\n", p.Name()))

			out, err := os.Create(fmt.Sprintf("%s.tar.gz.age", p.Name()))
			if err != nil {
				ui.Error(fmt.Errorf("cannot create file: %w", err).Error())
				os.Exit(genericExitCode)
			}
			defer out.Close()

			b, err := backup.NewTask(p.Name(), p.Location(), out)
			if err != nil {
				ui.Error(fmt.Sprintf("Cannot create backup task: %s", err))
				os.Exit(genericExitCode)
			}
			err = backup.Perform(b, passphrase)
			if err != nil {
				ui.Error(err.Error())
				os.Exit(genericExitCode)
			}

			ui.Output(fmt.Sprintf("Encryption passphrase is: %s", passphrase))
		} else {
			ui.Output("Not yet implemented")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// backupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	backupCmd.Flags().StringVar(&currentPersona, "persona", "", "The persona to backup")
}
