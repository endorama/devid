package manager

import (
	"fmt"
	"log"

	"github.com/endorama/devid/internal/plugin"
	"github.com/endorama/devid/internal/utils"
)

// DeregisterPlugin removes a plugin from the global plugin directory.
func DeregisterPlugin(plg plugin.Pluggable) error {
	pluginsDirectory[plg.Name()] = nil

	return nil
}

// RegisterPlugin register a plugin instance in the global plugin directory.
func RegisterPlugin(plg plugin.Pluggable, config []byte) error {
	pluginsDirectory[plg.Name()] = plg

	if configurablePlugin, ok := plg.(plugin.Configurable); ok {
		log.Printf("loading config for: %s", plg.Name())

		err := configurablePlugin.LoadConfig(config)
		if err != nil {
			return fmt.Errorf("loading plugin configuration failed: %w", err)
		}
	}

	return nil
}

// LoadCorePlugins instantiate and register all core plugins, configuring them using the values
// from the provided configuration file.
func LoadCorePlugins(configFile string) ([]error, error) {
	log.SetPrefix("core-plugins-loader ")
	defer log.SetPrefix("")

	errs := []error{}

	yamlFile, err := utils.ReadFile(configFile)
	if err != nil {
		return errs, fmt.Errorf("failed loading core plugins: %w", err)
	}

	for name, initFn := range plugin.Core {
		log.Printf("running for: %s", name)

		plg := initFn()

		err := RegisterPlugin(plg, yamlFile)
		if err != nil {
			errs = append(errs, err)
		}
	}

	return errs, nil
}

// LoadOptionalPlugins instantiate and register all optional plugins, configuring them using the
// values from the provided configuration file.
func LoadOptionalPlugins() ([]error, error) {
	log.SetPrefix("optional-plugins-loader ")
	defer log.SetPrefix("")

	errs := []error{}

	return errs, nil
}
