package aws

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	// LocalConfig specifies if AWS cli configuration should be overridden by a persona-specific
	// configuration. When true it the plugin will set AWS_CONFIG to a file within the persona
	// folder.
	// Default is false.
	LocalConfig       bool   `mapstructure:"local_config"`
	CustomProfileName string `mapstructure:"custom_profile_name"`
}

func (p *Plugin) Configure(v *viper.Viper) error {
	if err := v.Unmarshal(&p.config); err != nil {
		return fmt.Errorf("cannot unmarshal %s configuration:  %w", p.Name(), err)
	}

	return nil
}

func (p *Plugin) Config() Config {
	return p.config
}
