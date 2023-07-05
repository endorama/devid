package registry

import (
	"github.com/endorama/devid/internal/plugin"
	"github.com/endorama/devid/plugins/aws"
	"github.com/endorama/devid/plugins/ssh"
	"github.com/endorama/devid/plugins/tmux"
)

// Optionals are plugins that extend devid functionalities. Are loaded by devid after Core plugins.
func Optionals() []plugin.Pluggable {
	return []plugin.Pluggable{
		aws.NewPlugin(),
		ssh.NewPlugin(),
		tmux.NewPlugin(),
	}
}
