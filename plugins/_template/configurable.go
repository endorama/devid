package name

import (
	"fmt"

	"gopkg.in/yaml.v1"
)

type Config struct {
}

// Implements `plugin.Configurable` interface.
func (p Plugin) Config() interface{} {
	return p.config
}

// Implements `plugin.Configurable` interface.
func (p *Plugin) LoadConfig(configFile []byte) error {
	var config struct {
		Name struct {
			Config
		}
	}

	err := yaml.Unmarshal(configFile, &config)
	if err != nil {
		return fmt.Errorf("cannot unmarshal config file: %w", err)
	}

	p.config = config.Name.Config

	return nil
}
