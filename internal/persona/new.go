package persona

import (
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

	v.SetDefault("apiVersion", apiVersion)
	v.SetDefault("identity.name", "required")
	v.SetDefault("identity.email", "required")
	v.SetDefault("envs", map[string]string{})
	v.SetDefault("ssh.enabled", false)
	v.SetDefault("tmux.enabled", false)

	return Persona{
		location: loc,
		name:     name,
		Config:   v,
	}, nil
}
