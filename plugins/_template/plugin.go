package name

const pluginName = "template"

type Plugin struct {
	config Config
}

// NewPlugin instantiate a Plugin instance.
func NewPlugin() *Plugin {
	return &Plugin{}
}

// Name return plugin name.
// Implements `plugin.Pluggable` interface.
func (p Plugin) Name() string {
	return pluginName
}
