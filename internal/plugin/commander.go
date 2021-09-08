package plugin

import "github.com/spf13/cobra"

type Commander interface {
	Commands() []*cobra.Command
}
