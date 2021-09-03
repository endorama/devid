package envs

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config map[string]string

func (p *Plugin) Configure(v *viper.Viper) error {
	if err := v.Unmarshal(&p.config); err != nil {
		return fmt.Errorf("cannot unmarshal %s configuration:  %w", p.Name(), err)
	}

	return nil
}
