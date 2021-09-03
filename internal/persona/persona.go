package persona

import (
	"errors"
	"fmt"
	"path"

	"github.com/endorama/devid/internal/utils"
	"github.com/spf13/viper"
)

const (
	apiVersion = "v1"
	filename   = "config.yaml"
)

var errUnsupportedAPIVersion = errors.New("unsupported API version")

// Persona holds the entire persona information.
type Persona struct {
	location string
	name     string

	Config *viper.Viper
}

// Exists verify if a persona exists in the specified location
// For a persona to exists the configuration file should be present
// Persona configuration file is not validated.
func (p Persona) Exists() bool {
	ok := utils.Exists(p.location)
	if !ok {
		return false
	}

	ok = utils.Exists(p.File())

	return ok
}

func (p Persona) Load() error {
	err := p.Config.ReadInConfig()
	if err != nil {
		return fmt.Errorf("cannot read persona configuration: %w", err)
	}

	if p.Config.GetString("APIVersion") != apiVersion {
		return errUnsupportedAPIVersion
	}

	return nil
}

// File return profile configuration file path.
func (p Persona) File() string {
	return path.Join(p.location, filename)
}

// Location return persona's folder.
func (p Persona) Location() string {
	return p.location
}

// Name return persona's name.
func (p Persona) Name() string {
	return p.name
}

// Whoami returns human readable identity information.
func (p Persona) Whoami() string {
	return p.name
}
