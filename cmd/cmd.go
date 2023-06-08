package cmd

import (
	"github.com/endorama/devid/cmd/ui"
	"github.com/spf13/cobra"
)

var cli *cobra.Command

// Init initialises a cobra CLI with all commands from this package.
func Init() {
	cli = RootCmd()
	cli.AddCommand(
		Backup(),
		Delete(),
		Edit(),
		List(),
		New(),
		Rehash(),
		Shell(),
		Whoami(),
	)
}

// Execute perform execution of the global CLI initialised with Init().
// panics if Init() has not been called.
func Execute() {
	if cli == nil {
		panic("cli has not been initialised, have you called Init()?")
	}

	if err := cli.Execute(); err != nil {
		ui.Fatal(err, genericExitCode)
	}
}
