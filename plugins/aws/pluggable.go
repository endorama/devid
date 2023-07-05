package aws

// Name return plugin name.
// Implements `plugin.Pluggable` interface.
func (*Plugin) Name() string {
	return pluginName
}
