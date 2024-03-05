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
	"github.com/spf13/cobra"

	"github.com/endorama/devid/cmd/ui"
	"github.com/endorama/devid/internal/version"
)

// Version represents the version command.
func Version() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "print version",
		Long:  `Print version information bundled with the program.`,
		Run: func(_ *cobra.Command, _ []string) {
			ui.Outputf(version.BuildString())
		},
	}

	return versionCmd
}
