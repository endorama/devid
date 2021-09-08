package manager

import (
	"fmt"
	"log"

	"github.com/endorama/devid/internal/plugin"
	"github.com/endorama/devid/internal/plugin/registry"
	"github.com/spf13/cobra"
)

func LoadCommands() []*cobra.Command {
	log.SetPrefix("command-plugins-loader ")
	defer log.SetPrefix("")

	cmds := []*cobra.Command{}

	all := append(registry.Cores(), registry.Optionals()...) //nolint:gocritic // not appending to the same slice
	for _, plg := range all {
		log.Printf("%s plugin loading commands", plg.Name())

		if cmdPlg, ok := plg.(plugin.Commander); ok {
			short := fmt.Sprintf("devid %s", plg.Name())
			long := fmt.Sprintf(`%s

Provide access to %s dedicated subcommands.

This command does not do anything by itself, please use one of the available
subcommands.
`, short, plg.Name())

			wrapCmd := &cobra.Command{
				Use:   plg.Name(),
				Short: short,
				Long:  long,
			}
			wrapCmd.PersistentFlags().String("persona", "", "")

			for _, c := range cmdPlg.Commands() {
				log.Printf("adding: %s %s", plg.Name(), c.Name())
				wrapCmd.AddCommand(c)
			}

			cmds = append(cmds, wrapCmd)
		}
	}

	return cmds
}
