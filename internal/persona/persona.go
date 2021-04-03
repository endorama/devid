package persona

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"

	"github.com/endorama/devid/internal/plugin"
	"github.com/endorama/devid/internal/utils"
	"github.com/endorama/devid/plugins/identity"
	"gopkg.in/yaml.v1"
)

const (
	apiVersion = "v1"
	filename   = "config.yaml"
)

// Persona holds the entire persona information.
type Persona struct {
	// Version string allows to check for compatibility of the Persona configuration
	APIVersion string `yaml:"apiVersion"`

	location string `yaml:"-"`
	name     string `yaml:"-"`

	// Plugins contains a map of Pluggable
	Plugins map[string]plugin.Pluggable `yaml:",omitempty"`
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

func (p *Persona) Load() error {
	yamlFile, err := ioutil.ReadFile(p.File())
	if err != nil {
		return fmt.Errorf("cannot read yaml file: %w", err)
	}

	var fromConfig struct {
		APIVersion string   `yaml:"apiVersion"`
		Plugins    []string `yaml:"plugins"`
	}
	err = yaml.Unmarshal(yamlFile, &fromConfig)
	if err != nil {
		return fmt.Errorf("cannot unmarshal yaml file: %w", err)
	}

	if fromConfig.APIVersion != apiVersion {
		return fmt.Errorf("unsupported APIVersion")
	}

	p.APIVersion = fromConfig.APIVersion

	p.loadPlugins(fromConfig.Plugins)
	p.loadPluginsConfiguration(yamlFile)

	return nil
}

// Whoami returns human readable identity information.
func (p Persona) Whoami() string {
	return p.name
}

func (p Persona) loadPlugins(plugins []string) {
	for _, pluginName := range plugins {
		switch pluginName {
		case identity.PluginName:
			p.enablePlugin(identity.NewPlugin())
		default:
			panic(fmt.Sprintf("trying to load unknown plugin: %s", pluginName))
		}
	}
}

func (p *Persona) enablePlugin(pg plugin.Pluggable) {
	if p.Plugins[pg.Name()] == nil {
		p.Plugins[pg.Name()] = pg
	}
}

//
// func (p *Persona) disablePlugin(pg plugin.Pluggable) {
//   if p.Plugins[pg.Name()] != nil {
//     p.Plugins[pg.Name()] = nil
//   }
// }

func (p *Persona) loadPluginsConfiguration(configFile []byte) error {
	var errs []error

	for _, plg := range p.Plugins {
		if configurablePlugin, ok := plg.(plugin.Configurable); ok {
			err := configurablePlugin.LoadConfig(configFile)
			if err != nil {
				errs = append(errs, err)
			}
			log.Printf(fmt.Sprintf("  %+v\n", configurablePlugin.Config()))
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("error loading plugin configuration: %s", errs)
	}

	return nil
}
