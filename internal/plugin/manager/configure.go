package manager

import (
	"errors"
	"fmt"
	"log"

	"github.com/spf13/viper"

	"github.com/endorama/devid/internal/plugin"
)

// errMissingPluginConfigForName signal that a plugin with the provided name was not found.
var errMissingPluginConfigForName = errors.New("plugin config not found")

// configurePlugin is an helper that provides a plugin implementing the Configurable interface the
// relevant piece of the configuration (an object with the same name of the plugin is extracted from
// the configuration and provided to the plugin).
func configurePlugin(plg plugin.Pluggable, config *viper.Viper) error {
	if configurablePlugin, ok := plg.(plugin.Configurable); ok {
		log.Printf("configuring %s plugin", plg.Name())

		plgCfg := config.Sub(plg.Name())
		if plgCfg == nil {
			return fmt.Errorf("%s %w", plg.Name(), errMissingPluginConfigForName)
		}

		err := configurablePlugin.Configure(plgCfg)
		if err != nil {
			return fmt.Errorf("loading plugin configuration failed: %w", err)
		}
	}

	return nil
}
