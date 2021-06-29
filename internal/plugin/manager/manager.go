package manager

import (
	"fmt"
	"log"

	"github.com/endorama/devid/internal/plugin"
	"github.com/endorama/devid/internal/utils"
)

func DisablePlugin(plg plugin.Pluggable) error {
	enabledPlugins[plg.Name()] = nil

	return nil
}

func EnablePlugin(plg plugin.Pluggable, config []byte) error {
	enabledPlugins[plg.Name()] = plg

	if configurablePlugin, ok := plg.(plugin.Configurable); ok {
		err := configurablePlugin.LoadConfig(config)
		if err != nil {
			return fmt.Errorf("loading plugin configuration failed: %w", err)
		}
	}

	return nil
}

func LoadCorePlugins(configFile string) (error, []error) {
	log.SetPrefix("core-plugins-loader ")
	defer log.SetPrefix("")

	errs := []error{}

	yamlFile, err := utils.ReadFile(configFile)
	if err != nil {
		return fmt.Errorf("failed loading core plugins: %w", err), errs
	}

	for _, initFn := range corePlugins {
		plg := initFn()

		err := EnablePlugin(plg, yamlFile)
		if err != nil {
			errs = append(errs, err)
		}
	}

	return nil, errs
}

func LoadOptionalPlugins() (error, []error) {
	log.SetPrefix("optional-plugins-loader ")
	defer log.SetPrefix("")

	errs := []error{}

	return nil, errs
}
