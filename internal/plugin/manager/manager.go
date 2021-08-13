package manager

import (
	"fmt"
	"log"

	"github.com/endorama/devid/internal/plugin"
)

// DeregisterPlugin removes a plugin from the global plugin directory.
// func DeregisterPlugin(plg plugin.Pluggable) error {
//   pluginsDirectory[plg.Name()] = nil
//
//   return nil
// }

// RegisterPlugin register a plugin instance in the global plugin directory.
func RegisterPlugin(plg plugin.Pluggable, config plugin.Config, enabled bool) error {
	plugins = append(plugins, Plugin{Instance: plg, Enabled: enabled})

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

	for _, plg := range Core {
		name := plg.Name()
		log.Printf("running for: %s", name)

		// plg := initFn()

		err := RegisterPlugin(plg, config, true)
		if err != nil {
			errs = append(errs, err)
		}
	}

	return errs, nil
}

// LoadOptionalPlugins instantiate and register all optional plugins, configuring them using the
// values from the provided configuration file.
func LoadOptionalPlugins(config plugin.Config) ([]error, error) {
	log.SetPrefix("optional-plugins-loader ")
	defer log.SetPrefix("")

	errs := []error{}

	for _, plg := range Optional {
		enabled := false
		name := plg.Name()
		log.Printf("running for: %s", name)

		switch name {
		case "ssh":
			enabled = config.Ssh.Enabled
		case "tmux":
			enabled = config.Tmux.Enabled
		}

		err := RegisterPlugin(plg, config, enabled)
		if err != nil {
			errs = append(errs, err)
		}
	}

	return errs, nil
}

func LoadPlugins(config plugin.Config) ([]error, error) {
	errs, err := LoadCorePlugins(config)
	if err != nil {
		return errs, err
	}

	errs, err = LoadOptionalPlugins(config)
	if err != nil {
		return errs, err
	}

	return []error{}, nil
}

// GetPlugin find a plugin by name and return it. Boolean value varies if the
// plugin is found in the plugins list or not.
func GetPlugin(name string) (Plugin, bool) {
	for _, plg := range plugins {
		log.Printf("searching for: %s", plg.Instance.Name())

		if plg.Instance.Name() == name {
			return plg, true
		}
	}

	return Plugin{}, false
}
