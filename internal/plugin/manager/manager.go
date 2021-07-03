package manager

import (
	"fmt"
	"log"

	"github.com/endorama/devid/internal/plugin"
)

// DeregisterPlugin removes a plugin from the global plugin directory.
func DeregisterPlugin(plg plugin.Pluggable) error {
	pluginsDirectory[plg.Name()] = nil

	return nil
}

// RegisterPlugin register a plugin instance in the global plugin directory.
func RegisterPlugin(plg plugin.Pluggable, config plugin.Config) error {
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
func LoadCorePlugins(config plugin.Config) ([]error, error) {
	log.SetPrefix("core-plugins-loader ")
	defer log.SetPrefix("")

	errs := []error{}

	for name, initFn := range Core {
		log.Printf("running for: %s", name)

		plg := initFn()

		err := RegisterPlugin(plg, config)
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
