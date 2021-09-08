package registry

import (
	"github.com/endorama/devid/internal/plugin"
	"github.com/endorama/devid/plugins/ssh"
	"github.com/endorama/devid/plugins/tmux"
)

func Optionals() []plugin.Pluggable {
	return []plugin.Pluggable{
		ssh.NewPlugin(),
		tmux.NewPlugin(),
	}
}
