package manager

import (
	"github.com/endorama/devid/internal/plugin"
)

// enabledPlugins lists all plugins enabled in the current execution session
// nolint:gochecknoglobals // plugin are instantiated only once, so effectively "globals"
var enabledPlugins = make(map[string]plugin.Pluggable)

// Plugins returns the list of currently enabled plugins.
func Plugins() map[string]plugin.Pluggable {
	return enabledPlugins
}
