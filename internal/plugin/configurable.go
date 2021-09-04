package plugin

import "github.com/spf13/viper"

/*
Configurable interface allow a plugin to load configuration from the profile
folder.

A plugin implementing this interface will be able to use a part of the
Persona configuration file (the YAML object with the same name as the plugin
).
Plugin configuration happens before any other action is performed and
immediately after reading the Persona configuration file.
*/
type Configurable interface {
	Configure(v *viper.Viper) error
}
