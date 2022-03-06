package registry

import (
	"github.com/endorama/devid/internal/plugin"
	"github.com/endorama/devid/plugins/bin"
	"github.com/endorama/devid/plugins/envs"
	"github.com/endorama/devid/plugins/identity"
)

// Cores is a list of plugins that are core to devid working as intended. They cannot be disable
// and they MUST be initialized for the tool to work.
// This plugins can be relied upon by other plugins, as they will be present and initialed when
// optional plugins are leaded.
func Cores() []plugin.Pluggable {
	return []plugin.Pluggable{
		identity.NewPlugin(),
		bin.NewPlugin(),
		envs.NewPlugin(),
	}
}
