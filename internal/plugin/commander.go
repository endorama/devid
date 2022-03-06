package plugin

import "github.com/spf13/cobra"

/*
Commands interface allow a plugin to provide additional commands for devid CLI.

A plugin implementing this interface will be able to add commands to devid CLI.

NOTE that commands will be available independently from the enabled or disabled state of the plugin
in the user configuration, as command loading is performed before enablement status is evaluated.
*/
type Commander interface {
	Commands() []*cobra.Command
}
