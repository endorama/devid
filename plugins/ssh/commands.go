package ssh

import (
	"github.com/endorama/devid/plugins/ssh/cmds"
	"github.com/spf13/cobra"
)

func (p Plugin) Commands() []*cobra.Command {
	return []*cobra.Command{cmds.CreateKey, cmds.PrintPubKey}
}
