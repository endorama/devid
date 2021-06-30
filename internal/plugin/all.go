package plugin

import (
	"github.com/endorama/devid/plugins/bin"
	"github.com/endorama/devid/plugins/identity"
)

// PluggableInstantiator is a proxy type for the init function for a plugin.
type PluggableInstantiator func() Pluggable

// Core contains all core plugins. Core plugins are enabled by default and cannot be disabled.
// nolint:gochecknoglobals // needed for plugin discovery
var Core = map[string]PluggableInstantiator{
	"identity": func() Pluggable { return identity.NewPlugin() },
	"bin":      func() Pluggable { return bin.NewPlugin() },
}

// Optional contains all optional plugins. Optional plugins are disabled by default and can be enabled.
// nolint:gochecknoglobals // needed for plugin discovery
var Optional = map[string]PluggableInstantiator{}
