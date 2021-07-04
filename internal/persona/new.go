package persona

import (
	"errors"
	"fmt"
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

var errPersonaDoesNotExists = errors.New("does not exists")

func Load(name string) (Persona, error) {
	p, _ := New(name)

	if !p.Exists() {
		return p, errPersonaDoesNotExists
	}

	config, err := plugin.LoadConfigFromFile(p.File())
	if err != nil {
		return p, fmt.Errorf("cannot load configuration from file (%s): %w", p.File(), err)
	}

	p.Config = config

	return p, nil
}
