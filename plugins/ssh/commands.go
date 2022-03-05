package ssh

import (
	"github.com/spf13/cobra"

	"github.com/endorama/devid/plugins/ssh/cmds"
)

func (p Plugin) Commands() []*cobra.Command {
	return []*cobra.Command{cmds.CreateKey, cmds.PrintPubKey}
}
