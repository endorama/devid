package name

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	// TODO: add configuration fields
}

func (p *Plugin) Configure(v *viper.Viper) error {
	if err := v.Unmarshal(&p.config); err != nil {
		return fmt.Errorf("cannot unmarshal %s configuration:  %w", p.Name(), err)
	}

	// TODO: add config defaults
	// TODO: add config validation logic

	return nil
}
