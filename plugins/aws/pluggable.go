package aws

// Name return plugin name.
// Implements `plugin.Pluggable` interface.
func (_ *Plugin) Name() string {
	return pluginName
}
