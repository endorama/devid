package persona

import (
	"fmt"
	"path"

	"github.com/endorama/devid/internal/plugin"
	"github.com/endorama/devid/internal/utils"
)

const (
	apiVersion = "v1"
	filename   = "config.yaml"
)

// Persona holds the entire persona information
type Persona struct {
	// Version string allows to check for compatibility of the Persona configuration
	APIVersion string `yaml:"apiVersion"`

	location string `yaml:"-"`
	name     string `yaml:"-"`

	// Identity contains identifying information for this persona
	Identity Identity

	// Plugins contains a map of Pluggable
	Plugins map[string]plugin.Pluggable `yaml:",omitempty"`
}

// Exists verify if a persona exists in the specified location
// For a persona to exists the configuration file should be present
// Persona configuration file is not validated
func (p Persona) Exists() bool {
	ok := utils.Exists(p.location)
	if !ok {
		return false
	}
	ok = utils.Exists(p.File())
	return ok
}

// File return profile configuration file path
func (p Persona) File() string {
	return path.Join(p.location, filename)
}

// Location return persona's folder
func (p Persona) Location() string {
	return p.location
}

// Name return persona's name
func (p Persona) Name() string {
	return p.name
}

func (p Persona) Load() error {
	return nil
}

// Whoami returns human readable identity information
func (p Persona) Whoami() string {
	return fmt.Sprintf("%s <%s>", p.Identity.Name, p.Identity.Email)
}

func (p Persona) loadPlugins() {}

func (p *Persona) enablePlugin(pg plugin.Pluggable) {
	if p.Plugins[pg.Name()] == nil {
		p.Plugins[pg.Name()] = pg
	}
}

func (p *Persona) disablePlugin(pg plugin.Pluggable) {
	if p.Plugins[pg.Name()] != nil {
		p.Plugins[pg.Name()] = nil
	}
}
