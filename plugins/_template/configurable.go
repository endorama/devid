package name

import "github.com/endorama/devid/internal/plugin"

type Config struct {
	// TODO: add configuration fields
	// TODO: add configuration fields to plugin.Config struct
}

func (p Plugin) Config() interface{} {
	return p.config
}

func (p *Plugin) LoadConfig(config plugin.Config) error {
	// TODO: import configuration option from plugin.Config
	// TODO: add key for loading from external config in internal/plugin/from-file.go LoadConfig()
	// p.config = config....
	return nil
}
