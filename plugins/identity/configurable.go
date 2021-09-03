package identity

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Email string `yaml:"email"`
	Name  string `yaml:"name"`
}

var errEmailMissing = errors.New("email field missing")
var errNameMissing = errors.New("name field missing")

func (p *Plugin) Configure(v *viper.Viper) error {
	if err := v.Unmarshal(&p.config); err != nil {
		return fmt.Errorf("cannot unmarshal %s configuration:  %w", p.Name(), err)
	}

	if p.config.Email == "" {
		return errEmailMissing
	}

	if p.config.Name == "" {
		return errNameMissing
	}

	return nil
}
