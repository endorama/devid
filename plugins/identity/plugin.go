package identity

import (
	"fmt"
)

const PluginName = "identity"

func NewPlugin() *Plugin {
	return &Plugin{}
}

type Plugin struct {
	config Config
}

func (p Plugin) Name() string {
	return PluginName
}

func (p Plugin) Whoami() string {
	cfg := p.config

	return fmt.Sprintf("%s <%s>", cfg.Name, cfg.Email)
}
