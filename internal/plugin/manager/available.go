package manager

import (
	"github.com/endorama/devid/internal/plugin"
)

type Plugin struct {
	Instance plugin.Pluggable
	Enabled  bool
}

var plugins = []Plugin{}
