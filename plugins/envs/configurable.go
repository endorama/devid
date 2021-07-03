package envs

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Envs map[string]string
}

func (p Plugin) Config() interface{} {
	return p.config
}

func (p *Plugin) LoadConfig(configFile []byte) error {
	var config struct {
		Envs struct {
			Config
		}
	}

	err := yaml.Unmarshal(configFile, &config)
	if err != nil {
		return fmt.Errorf("cannot unmarshal config file: %w", err)
	}

	p.config = config.Envs.Config

	return nil
}
