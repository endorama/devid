package plugin

import "github.com/endorama/devid/plugins/identity"

const (
	apiVersion = "v1"
)

func NewConfig() Config {
	return Config{
		APIVersion: apiVersion,
	}
}

type Config struct {
	APIVersion string `yaml:"apiVersion"`

	Identity struct {
		identity.Config
	}
}
