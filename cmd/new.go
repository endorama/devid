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
// nolint:gochecknoglobals
package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/endorama/devid/internal/persona"
	"github.com/endorama/devid/internal/utils"
)

// newCmd represents the new command.
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new (empty) persona",
	Long: `devid new <persona name>

Create a new persona configuration file, opens it within EDITOR.
`,
	Run: func(cmd *cobra.Command, args []string) {
		runCommand(args)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
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

	err = utils.OpenWithEditor(p.File())
	if err != nil {
		ui.Error(err.Error())
		os.Exit(genericExitCode)
	}
}
