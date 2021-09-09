package manager

import (
	"fmt"

	"github.com/endorama/devid/internal/plugin"
	"github.com/endorama/devid/internal/plugin/registry"
	"github.com/spf13/cobra"
)

func LoadCommands() []*cobra.Command {
	// loggin here cannot be managed via --verbose flag, so disabling it
	// log.SetPrefix("command-plugins-loader ")
	// defer log.SetPrefix("")

	cmds := []*cobra.Command{}

	all := append(registry.Cores(), registry.Optionals()...) //nolint:gocritic // not appending to the same slice
	for _, plg := range all {
		// log.Printf("%s plugin loading commands", plg.Name())

		if cmdPlg, ok := plg.(plugin.Commander); ok {
			wrapCmd := &cobra.Command{
				Use:   plg.Name(),
				Short: fmt.Sprintf("%s plugin subcommands", plg.Name()),
				Long: fmt.Sprintf(`Provide access to %s dedicated subcommands.

This command does not do anything by itself, please use one of the available
subcommands.
`, plg.Name()),
			}
			wrapCmd.PersistentFlags().String("persona", "", "persona on which to execute the specified action")

			for _, c := range cmdPlg.Commands() {
				// log.Printf("adding: %s %s", plg.Name(), c.Name())
				wrapCmd.AddCommand(c)
			}

			cmds = append(cmds, wrapCmd)
		}
	}

	return cmds
}
