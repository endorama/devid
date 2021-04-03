package identity

import (
	"fmt"

	"gopkg.in/yaml.v1"
)

type Config struct {
	Email string `yaml:"email"`
	Name  string `yaml:"name"`
}

func (p Plugin) Config() interface{} {
	return p.config
}

func (p *Plugin) LoadConfig(configFile []byte) error {
	var config struct {
		Identity struct {
			Email string `yaml:"email"`
			Name  string `yaml:"name"`
		}
	}

	err := yaml.Unmarshal(configFile, &config)
	if err != nil {
		return fmt.Errorf("cannot unmarshal config file: %w", err)
	}

	p.config = Config{
		Email: config.Identity.Email,
		Name:  config.Identity.Name,
	}

	return nil
}
