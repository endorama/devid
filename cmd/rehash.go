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
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/endorama/devid/cmd/ui"
	cmdutils "github.com/endorama/devid/cmd/utils"
	"github.com/endorama/devid/internal/persona"
	"github.com/endorama/devid/internal/plugin/manager"
	"github.com/endorama/devid/internal/utils"
)

const permUserRWX = os.FileMode(0700)

// rehashCmd represents the rehash command.
func Rehash() *cobra.Command {
	rehashCmd := &cobra.Command{ //nolint:gochecknoglobals // required by cobra
		Use:   "rehash",
		Short: "rebuild profiles loader and shims",
		Long: `Rebuild profiles loader and shims, performing setup, generation and shell file rendering.

This command will not run if a personal is already loaded.

At it's core devid is a tool to render a bash based template to load a dedicated environment for a
specific persona.

This resulting file is stored in the persona's folder and is usable without devid. If everything 
goes bad you can still load a persona's shell environment by running 
	bash /path/to/persona/load.sh

rehash command is directly inspired by rbenv, a ruby version manager.
`,
		Run: runRehash,
	}

	rehashCmd.Flags().String("persona", "", "The persona for which to rebuild the shell configuration")

	return rehashCmd
}

func runRehash(cmd *cobra.Command, _ []string) {
	if viper.GetString("active_persona") != "" {
		// NOTE: rehashing when a profile is active is dangerous, as the environment
		// has been changed with customizations and there is no guarantee about
		// what those changes have affected.
		// This may be especially problematic for executable path detection in
		// plugin.
		// As such we prevent rehashing while there is an active profile.
		ui.Fatal(errRehashWithActiveProfile, genericExitCode)
	}

	var (
		errs []error
		err  error
	)

	p, err := cmdutils.LoadPersona(cmd)
	if err != nil {
		ui.Fatal(fmt.Errorf("cannot instantiate persona: %w", err), noPersonaLoadedExitCode)
	}
	// errs, err := manager.LoadPlugins(p.Config)
	// if err != nil {
	//   ui.Error(err.Error())
	//
	//   for _, e := range errs {
	//     ui.Error(e.Error())
	//   }
	//
	//   os.Exit(pluginManagerLoadingErrorExitCode)
	// }

	errs, err = manager.LoadCorePlugins(p.Config)
	if err != nil {
		handleAllErrors(errs, err, pluginManagerCoreLoadingErrorExitCode)
	}

	errs, err = manager.LoadOptionalPlugins(p.Config)
	if err != nil {
		handleAllErrors(errs, err, pluginManagerOptionalLoadingErrorExitCode)
	}

	log.Printf("persona: %+v\n", p)

	errs, err = manager.SetupPlugins(p)
	if err != nil {
		handleAllErrors(errs, err, pluginGenerationExitCode)
	}

	errs, err = manager.Generate(p)
	if err != nil {
		handleAllErrors(errs, err, pluginGenerationExitCode)
	}

	fn := viper.GetString("shell_loader_filename")
	if err := writeShellLoader(p, fn); err != nil {
		ui.Fatal(err, genericExitCode)
	}
}

func handleAllErrors(errs []error, err error, exitCode int) {
	ui.Error(err)

	for _, e := range errs {
		ui.Error(e)
	}

	os.Exit(exitCode)
}

func writeShellLoader(p persona.Persona, filename string) error {
	content, err := manager.ShellLoader(p)
	if err != nil {
		ui.Error(err)
		os.Exit(1)
	}

	log.Printf("%+v\n", content)

	shellLoaderFilePath := path.Join(p.Location(), filename)

	err = utils.PersistFile(shellLoaderFilePath, content)
	if err != nil {
		return fmt.Errorf("cannot save shell loader: %w", err)
	}

	if err := os.Chmod(shellLoaderFilePath, permUserRWX); err != nil {
		return fmt.Errorf("cannot change permissions to %s on %s: %w", permUserRWX, shellLoaderFilePath, err)
	}

	return nil
}
