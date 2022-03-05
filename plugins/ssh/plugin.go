package ssh

import "github.com/spf13/afero"

const pluginName = "ssh"

type Plugin struct {
	config Config
	fs     afero.Fs
}

// NewPlugin instantiate a Plugin instance.
func NewPlugin() *Plugin {
	return &Plugin{
		fs: afero.NewOsFs(),
	}
}
