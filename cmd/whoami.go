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
	"github.com/endorama/devid/internal/plugin/manager"
	"github.com/endorama/devid/plugins/identity"
)

// whoamiCmd represents the whoami command.
func Whoami() *cobra.Command {
	whoamiCmd := &cobra.Command{ //nolint:gochecknoglobals // required by cobra
		Use:   "whoami",
		Short: "print current loaded persona",
		Long: `Print information about persona currently loaded.

If no persona is loaded print nothing and exit with code 128.

This command loads the current persona from DEVID_ACTIVE_PERSONA environment variable, and this 
value takes precedence over the --persona flag.
`,
		Run: func(cmd *cobra.Command, args []string) {
			p, err := utils.LoadPersona(cmd)
			if err != nil {
				ui.Fatal(fmt.Errorf("cannot instantiate persona: %w", err), noPersonaLoadedExitCode)
			}

			extended, err := cmd.Flags().GetBool("extended")
			if err != nil {
				ui.Error(fmt.Errorf("cannot access flag extended: %w", err))
			}

			if extended {
				if errs, err := manager.LoadCorePlugins(p.Config); err != nil {
					ui.Error(errLoadingCorePlugins)

					for _, e := range errs {
						ui.Error(e)
					}

					os.Exit(pluginManagerCoreLoadingErrorExitCode)
				}

				plg, found := manager.GetPlugin("identity")
				if !found {
					ui.Fatal(errIdentityPluginNotFound, genericExitCode)
				}

				identityPlugin, ok := plg.Instance.(*identity.Plugin)
				if !ok {
					ui.Fatal(errPluginNotInstanceOf, genericExitCode)
				}

				ui.Outputf("%s, %s", p.Name(), identityPlugin.Whoami())
				os.Exit(0)
			}

			ui.Outputf(p.Name())
			os.Exit(0)
		},
	}

	whoamiCmd.Flags().BoolP("extended", "e", false, "Print extended identity information")
	return whoamiCmd
}
