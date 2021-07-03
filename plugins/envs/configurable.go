package envs

import "github.com/endorama/devid/internal/plugin"

type Config map[string]string

func (p Plugin) Config() interface{} {
	return p.config
}

func (p *Plugin) LoadConfig(config plugin.Config) error {
	p.config = config.Envs
	return nil
}
