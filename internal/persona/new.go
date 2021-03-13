package persona

import (
	"path"

	"github.com/spf13/viper"
)

func New(name string) (Persona, error) {
	location := viper.GetString("personas_location")

	return Persona{
		APIVersion: apiVersion,
		location: path.Join(location, name),
		name: name,
	}, nil
}

func NewWithCustomLocation(name, location string) (Persona, error) {
	return Persona{
		APIVersion: apiVersion,
		location: path.Join(location, name),
		name: name,
	}, nil
}

