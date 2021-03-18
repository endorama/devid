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
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/endorama/devid/internal/persona"
)

// listCmd represents the list command.
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available personas",
	Long: `devid list

List all available personas.
`,
	Run: func(cmd *cobra.Command, args []string) {
		files, err := ioutil.ReadDir(viper.GetString("personas_location"))
		if err != nil {
			ui.Error(fmt.Errorf("cannot read folder content: %w", err).Error())
			os.Exit(genericExitCode)
		}
		for _, f := range files {
			if f.IsDir() {
				p, _ := persona.New(f.Name())
				if p.Exists() {
					ui.Output(p.Name())
				}
			}

		}
	},
}

func init() { //nolint:gochecknoinits // required by cobra
	rootCmd.AddCommand(listCmd)
}
