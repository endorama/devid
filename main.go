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
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"

	"github.com/endorama/devid/cmd"
	"github.com/endorama/devid/internal/settings"
)

func main() {
	settings.Init()

	createPersonasFolder()

	rootCmd := cmd.RootCmd()
	rootCmd.AddCommand(cmd.Backup())
	rootCmd.AddCommand(cmd.Delete())
	rootCmd.AddCommand(cmd.Edit())
	rootCmd.AddCommand(cmd.List())
	rootCmd.AddCommand(cmd.New())
	rootCmd.AddCommand(cmd.Rehash())
	rootCmd.AddCommand(cmd.Shell())
	rootCmd.AddCommand(cmd.Whoami())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func createPersonasFolder() {
	loc := viper.GetString("personas_location")

	const perm = os.FileMode(0750)

	if _, err := os.Stat(loc); os.IsNotExist(err) {
		log.Printf("%s does not exists, creating\n", loc)

		if err := os.MkdirAll(loc, perm); err != nil {
			log.Fatal(err)
		}
	}
}
