package name

// TODO: add name.NewPlugin() to internal/plugin/manager/available.go under Optional
// TODO: add enabled check in internal/plugin/manager/manager.go LoadOptionalPlugins switch case
const pluginName = "name"

type Plugin struct {
	config Config
}

// NewPlugin instantiate a Plugin instance.
func NewPlugin() *Plugin {
	return &Plugin{}
}
