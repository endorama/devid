package identity

import "github.com/endorama/devid/internal/plugin"

type Config struct {
	Email string `yaml:"email"`
	Name  string `yaml:"name"`
}

func (p Plugin) Config() interface{} {
	return p.config
}

func (p *Plugin) LoadConfig(config plugin.Config) error {
	p.config = Config{
		Email: config.Identity.Email,
		Name:  config.Identity.Name,
	}

	return nil
}
