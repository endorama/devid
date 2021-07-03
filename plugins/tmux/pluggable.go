package tmux

// Name return plugin name.
// Implements `plugin.Pluggable` interface.
func (p Plugin) Name() string {
	return pluginName
}
