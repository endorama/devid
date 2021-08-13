package manager

import (
	"github.com/endorama/devid/internal/plugin"
	"github.com/endorama/devid/plugins/bin"
	"github.com/endorama/devid/plugins/envs"
	"github.com/endorama/devid/plugins/identity"
	"github.com/endorama/devid/plugins/ssh"
	"github.com/endorama/devid/plugins/tmux"
)

// pluginsDirectory lists all plugins enabled in the current execution session
// nolint:gochecknoglobals // plugin are instantiated only once, so effectively "globals"
// var pluginsDirectory = make(map[string]plugin.Pluggable)
//
// // Plugins returns the list of currently enabled plugins.
// func Plugins() map[string]plugin.Pluggable {
//   return pluginsDirectory
// }

// Core contains all core plugins. Core plugins are enabled by default and cannot be disabled.
// nolint:gochecknoglobals // needed for plugin discovery
// var Core = map[string]plugin.PluggableInstantiator{
//   "identity": func() plugin.Pluggable { return identity.NewPlugin() },
//   "bin":      func() plugin.Pluggable { return bin.NewPlugin() },
//   "envs":     func() plugin.Pluggable { return envs.NewPlugin() },
// }

var Core = []plugin.Pluggable{
	identity.NewPlugin(),
	bin.NewPlugin(),
	envs.NewPlugin(),
}

// Optional contains all optional plugins. Optional plugins are disabled by default and can be enabled.
// nolint:gochecknoglobals // needed for plugin discovery
// var Optional = map[string]plugin.PluggableInstantiator{
//   "tmux": func() plugin.Pluggable { return tmux.NewPlugin() },
// }
var Optional = []plugin.Pluggable{
	ssh.NewPlugin(),
	tmux.NewPlugin(),
}

type Plugin struct {
	Instance plugin.Pluggable
	// Instantiator plugin.PluggableInstantiator
	Enabled bool
}

var plugins = []Plugin{}
