package persona

import (
	"path"

	"github.com/spf13/viper"

	"github.com/endorama/devid/internal/plugin"
)

func New(name string) (Persona, error) {
	location := viper.GetString("personas_location")

	return NewWithCustomLocation(name, location)
}

func NewWithCustomLocation(name, location string) (Persona, error) {
	return Persona{
		location: path.Join(location, name),
		name:     name,
		Config:   plugin.NewConfig(),
	}, nil
}
