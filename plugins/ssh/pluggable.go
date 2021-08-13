package ssh

// Name return plugin ssh.
// Implements `plugin.Pluggable` interface.
func (p Plugin) Name() string {
	return pluginName
}
