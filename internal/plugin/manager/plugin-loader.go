package manager

import (
	"errors"
	"fmt"
	"log"

	"github.com/endorama/devid/internal/plugin/registry"

	"github.com/spf13/viper"
)

var errLoadingCorePlugins = errors.New("cannot load all core plugins")
var errLoadingOptionalPlugins = errors.New("cannot load all requested optional plugins")

// LoadCorePlugins instantiate and register all core plugins, configuring them using the values
// from the provided configuration file.
func LoadCorePlugins(config *viper.Viper) ([]error, error) {
	log.SetPrefix("core-plugins-loader ")
	defer log.SetPrefix("")

	errs := []error{}

	for _, plg := range registry.Cores() {
		name := plg.Name()
		log.Printf("running for: %s", name)

		err := configurePlugin(plg, config)
		if err != nil {
			errs = append(errs, err)

			continue
		}

		plugins = append(plugins, Plugin{Instance: plg, Enabled: true})
	}

	if len(errs) != 0 {
		return errs, errLoadingCorePlugins
	}

	return errs, nil
}

// LoadOptionalPlugins instantiate and register all optional plugins, configuring them using the
// values from the provided configuration file.
func LoadOptionalPlugins(config *viper.Viper) ([]error, error) {
	log.SetPrefix("optional-plugins-loader ")
	defer log.SetPrefix("")

	errs := []error{}

	for _, plg := range registry.Optionals() {
		name := plg.Name()
		log.Printf("running for: %s", name)

		enabled := config.GetBool(fmt.Sprintf("%s.enabled", name))

		err := configurePlugin(plg, config)
		if err != nil && !errors.Is(err, errMissingPluginConfig) {
			errs = append(errs, err)

			continue
		}

		if err != nil && errors.Is(err, errMissingPluginConfig) {
			enabled = false
		}

		log.Printf("%s plugin is %s", name, humanizeEnabled(enabled))

		plugins = append(plugins, Plugin{Instance: plg, Enabled: enabled})
	}

	if len(errs) != 0 {
		return errs, errLoadingOptionalPlugins
	}

	return errs, nil
}

func LoadPlugins(config *viper.Viper) ([]error, error) {
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

func humanizeEnabled(enabled bool) string {
	if enabled {
		return "enabled"
	}

	return "disabled"
}
