package ssh

import "github.com/endorama/devid/internal/plugin"

const defaultCachePath = "/tmp/devid-%s-ssh-agent.tmp"

type Config struct {
	Keys      []string
	CachePath string
}

func (p Plugin) Config() interface{} {
	return p.config
}

func (p *Plugin) LoadConfig(config plugin.Config) error {
	p.config.Keys = config.Ssh.Keys

	p.config.CachePath = config.Ssh.CachePath
	if p.config.CachePath == "" {
		p.config.CachePath = defaultCachePath
	}

	return nil
}
