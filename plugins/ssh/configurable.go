package ssh

import (
	"fmt"

	"github.com/spf13/viper"
)

const defaultCachePath = "/tmp/devid-%s-ssh-agent.tmp"

var defaultSshKeys = []string{"id_rsa"}

type Config struct {
	Keys      []string
	CachePath string
}

func (p *Plugin) Configure(v *viper.Viper) error {
	if err := v.Unmarshal(&p.config); err != nil {
		return fmt.Errorf("cannot unmarshal %s configuration:  %w", p.Name(), err)
	}

	if p.config.CachePath == "" {
		p.config.CachePath = defaultCachePath
	}

	if len(p.config.Keys) == 0 {
		p.config.Keys = defaultSshKeys
	}

	return nil
}
