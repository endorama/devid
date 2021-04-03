package loader

import (
	"errors"
	"fmt"
	"log"

	"github.com/endorama/devid/internal/persona"
	"github.com/endorama/devid/internal/plugin"
)

var errLoadingPlugins = errors.New("error loading plugins configuration")

func loadPluginsConfiguration(p persona.Persona, configFile []byte) error {
	var errs []error

	for _, plg := range p.Plugins {
		if configurablePlugin, ok := plg.(plugin.Configurable); ok {
			err := configurablePlugin.LoadConfig(configFile)
			if err != nil {
				errs = append(errs, err)
			}

			log.Printf("  %+v\n", configurablePlugin.Config())
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("%w: %s", errLoadingPlugins, errs)
	}

	return nil
}
