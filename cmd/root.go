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
	"github.com/endorama/devid/cmd/ui"
	"github.com/endorama/devid/internal/plugin/manager"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string //nolint:gochecknoglobals // required for init

// var ui = &cli.ColoredUi{ //nolint:gochecknoglobals // UI is shared
//   OutputColor: cli.UiColorNone,
//   InfoColor:   cli.UiColorNone,
//   ErrorColor:  cli.UiColorRed,
//   WarnColor:   cli.UiColorYellow,
//
//   Ui: &cli.BasicUi{
//     Reader:      os.Stdin,
//     Writer:      os.Stdout,
//     ErrorWriter: os.Stderr,
//   },
// }

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{ //nolint:gochecknoglobals // required by cobra
	Use:   "devid",
	Short: "Secure manager for your developer personas",
	Long: `devid (pronounced /ˈdeɪvɪd/) is a Swiss Army Knife for your developer identity personas.

Each of us has multiple personas for different areas of their life. It may be work/personal, or for different open source projects, for different clients, or whatever reason you may think for presenting yourself differently in different context. This is something we do in real life (think dressing differently for different social events) but doing so in digital world as developers can be a pain: you have to manage identities (GPG or SSH keys), authentication tokens, specific configurations.

Properly securing our developer identity and personas is hard. devid aims to help you with that.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		ui.Fatal(err, genericExitCode)
	}
}

func init() { //nolint:gochecknoinits // required by cobra
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.devid.yaml)")

	rootCmd.AddCommand(manager.LoadCommands()...)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			ui.Fatal(err, genericExitCode)
		}

		// Search config in home directory with name ".devid" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".devid")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		ui.Info("Using config file: %s", viper.ConfigFileUsed())
	}
}
