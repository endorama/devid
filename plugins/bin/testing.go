package bin

import "github.com/spf13/afero"

// TestNewPlugin return a test utility function to instantiate a plugin instance
// with the specified afero.Fs instance.
func TestNewPlugin(fs afero.Fs) *Plugin {
	return &Plugin{
		fs: fs,
	}
}
