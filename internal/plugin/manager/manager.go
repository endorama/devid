package manager

import (
	"log"

	"github.com/endorama/devid/internal/plugin"
)

// Plugins is a managed plugin instance with information about it's enabled or disabled state.
type Plugin struct {
	Instance plugin.Pluggable
	Enabled  bool
}

//nolint:gochecknoglobals // the plugins state is shared through this global variable
var plugins = []Plugin{}

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
