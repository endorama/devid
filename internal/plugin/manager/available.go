package manager

import (
	"github.com/endorama/devid/internal/plugin"
	"github.com/endorama/devid/plugins/identity"
)

type PluggableInstantiator func() *identity.Plugin

var enabledPlugins = make(map[string]plugin.Pluggable)

var corePlugins = map[string]PluggableInstantiator{
	"identity": identity.NewPlugin,
}

var optionalPlugins = map[string]PluggableInstantiator{}

func Plugins() map[string]plugin.Pluggable {
	return enabledPlugins
}
