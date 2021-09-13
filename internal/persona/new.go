package persona

import (
	"errors"
	"fmt"
	"path"

	"github.com/spf13/viper"
)

func New(name string) (Persona, error) {
	location := viper.GetString("personas_location")

	return NewWithCustomLocation(name, location)
}

func NewWithCustomLocation(name, location string) (Persona, error) {
	loc := path.Join(location, name)

	v := viper.New()
	v.SetConfigType(configType)
	v.SetConfigName(filename)
	v.AddConfigPath(loc)

	return Persona{
		location: loc,
		name:     name,
		Config:   v,
	}, nil
}

var errPersonaDoesNotExists = errors.New("does not exists")

func Load(name string) (Persona, error) {
	p, err := New(name)
	if err != nil {
		return p, fmt.Errorf("init failed: %w", err)
	}

	if !p.Exists() {
		return p, fmt.Errorf("%w in %s", errPersonaDoesNotExists, p.Location())
	}

	err = p.Load()
	if err != nil {
		return p, fmt.Errorf("cannot load persona configuration: %w", err)
	}

	return p, nil
}
