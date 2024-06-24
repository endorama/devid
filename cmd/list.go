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
	"github.com/spf13/viper"

	"github.com/endorama/devid/cmd/ui"
	"github.com/endorama/devid/internal/persona"
)

// listCmd represents the list command.
func List() *cobra.Command {
	listCmd := &cobra.Command{ //nolint:gochecknoglobals // required by cobra
		Use:   "list",
		Short: "list personas",
		Long:  `List all available personas.`,
		Run: func(_ *cobra.Command, _ []string) {
			files, err := os.ReadDir(viper.GetString("personas_location"))
			if err != nil {
				ui.Fatal(fmt.Errorf("cannot read folder content: %w", err), noPersonaLoadedExitCode)
			}
			for _, f := range files {
				if f.IsDir() {
					p, _ := persona.New(f.Name())
					if p.Exists() {
						ui.Outputf(p.Name())
					}
				}
			}
		},
	}

	return listCmd
}
