package manager

import (
	"errors"
	"fmt"
	"log"

	"github.com/spf13/viper"

	"github.com/endorama/devid/internal/plugin"
)

var errMissingPluginConfig = errors.New("plugin config not found")

func configurePlugin(plg plugin.Pluggable, config *viper.Viper) error {
	if configurablePlugin, ok := plg.(plugin.Configurable); ok {
		log.Printf("configuring %s plugin", plg.Name())

		plgCfg := config.Sub(plg.Name())
		if plgCfg == nil {
			return fmt.Errorf("%s %w", plg.Name(), errMissingPluginConfig)
		}

		err := configurablePlugin.Configure(plgCfg)
		if err != nil {
			return fmt.Errorf("loading plugin configuration failed: %w", err)
		}
	}

	return nil
}
