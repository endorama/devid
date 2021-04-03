package persona

import (
	"path"

	"github.com/endorama/devid/internal/plugin"
	"github.com/spf13/viper"
)

func New(name string) (Persona, error) {
	location := viper.GetString("personas_location")

	return Persona{
		APIVersion: apiVersion,
		location:   path.Join(location, name),
		name:       name,
		Plugins:    make(map[string]plugin.Pluggable),
	}, nil
}

func NewWithCustomLocation(name, location string) (Persona, error) {
	return Persona{
		APIVersion: apiVersion,
		location:   path.Join(location, name),
		name:       name,
		Plugins:    make(map[string]plugin.Pluggable),
	}, nil
}
