package name

const pluginName = "name"

type Plugin struct {
	config Config
}

// NewPlugin instantiate a Plugin instance.
func NewPlugin() *Plugin {
	return &Plugin{}
}
