package gcp

// TODO: add gcp.NewPlugin() to internal/plugin/manager/available.go under Optional.
// TODO: add enabled check in internal/plugin/manager/manager.go LoadOptionalPlugins switch case.
const pluginName = "gcp"

type Plugin struct {
}

// NewPlugin instantiate a Plugin instance.
func NewPlugin() *Plugin {
	return &Plugin{}
}
