package gcp

// Name return plugin gcp.
// Implements `plugin.Pluggable` interface.
func (p Plugin) Name() string {
	return pluginName
}
