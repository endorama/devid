package registry

import (
	"github.com/endorama/devid/internal/plugin"
	"github.com/endorama/devid/plugins/bin"
	"github.com/endorama/devid/plugins/envs"
	"github.com/endorama/devid/plugins/identity"
)

func Cores() []plugin.Pluggable {
	return []plugin.Pluggable{
		identity.NewPlugin(),
		bin.NewPlugin(),
		envs.NewPlugin(),
	}
}
