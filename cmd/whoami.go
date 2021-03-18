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

	"github.com/endorama/devid/internal/persona"
)

var (
	extended bool
)

// whoamiCmd represents the whoami command.
var whoamiCmd = &cobra.Command{ //nolint:gochecknoglobals // required by cobra
	Use:   "whoami",
	Short: "Print current loaded persona",
	Long: `Print current loaded persona.

If no persona is loaded print nothing and exit with code 128.
`,
	Run: func(cmd *cobra.Command, args []string) {
		currentProfile := os.Getenv("DEVENV_ACTIVE_PROFILE")

		// there is no loaded profile
		if currentProfile == "" {
			os.Exit(noPersonalLoadedExitCode)
		}

		if extended {
			p, _ := persona.New(currentProfile)
			currentProfile = fmt.Sprintf("%s (%s)", currentProfile, p.Whoami())
		}

		ui.Output(currentProfile)
		os.Exit(0)
	},
}

func init() { //nolint:gochecknoinits // required by cobra
	rootCmd.AddCommand(whoamiCmd)

	whoamiCmd.Flags().BoolVarP(&extended, "extended", "e", false, "Print extended identity information")
}
