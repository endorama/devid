package manager

import (
	"fmt"
	"log"

	"github.com/endorama/devid/internal/persona"
	"github.com/endorama/devid/internal/plugin"
)

// setupPlugin execute Setup function a plugin implementing Setuppable interface.
func setupPlugin(p persona.Persona, plg plugin.Pluggable) error {
	if setuppablePlugin, ok := plg.(plugin.Setuppable); ok {
		log.Printf("running setup of: %s", plg.Name())

		err := setuppablePlugin.Setup(p.Location())
		if err != nil {
			return fmt.Errorf("plugin setup failed: %w", err)
		}
	}

	return nil
}

// SetupPlugins perform setup for all initialized plugins implementing Setuppable interface.
func SetupPlugins(p persona.Persona) ([]error, error) {
	log.SetPrefix("plugins-setup ")
	defer log.SetPrefix("")

	errs := []error{}

	for _, plg := range plugins {
		err := setupPlugin(p, plg.Instance)
		if err != nil {
			errs = append(errs, err)
		}
	}

	return errs, nil
}
