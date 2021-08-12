package bin

import "github.com/spf13/afero"

const pluginName = "bin"

type Plugin struct {
	fs afero.Fs
}

func NewPlugin() *Plugin {
	return &Plugin{
		fs: afero.NewOsFs(),
	}
}
