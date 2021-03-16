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
	"os"

	"github.com/endorama/devid/internal/persona"
	"github.com/endorama/devid/internal/utils"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a persona definition file in your $EDITOR",
	Long: `devid edit <persona name>

Open within EDITOR the specified persona configuration file.
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			ui.Error("Argument NAME required")
			os.Exit(genericExitCode)
		}

		name := args[0]

		p, err := persona.New(name)
		if err != nil {
			ui.Error(err.Error())
			os.Exit(genericExitCode)
		}

		if !p.Exists() {
			ui.Error("Persona does not exists")
			os.Exit(genericExitCode)
		}

		err = utils.OpenWithEditor(p.File())
		if err != nil {
			ui.Error(err.Error())
			os.Exit(genericExitCode)
		}
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
