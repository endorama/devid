package loader

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v1"

	"github.com/endorama/devid/internal/persona"
	"github.com/endorama/devid/internal/plugin"
	"github.com/endorama/devid/plugins/identity"
)

func LoadPlugins(p *persona.Persona) error {
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

	if fromConfig.APIVersion != p.APIVersion {
		return fmt.Errorf("unsupported APIVersion")
	}

	p.APIVersion = fromConfig.APIVersion

	loadPlugins(*p, fromConfig.Plugins)
	loadPluginsConfiguration(*p, yamlFile)

	return nil

}

func loadPlugins(p persona.Persona, plugins []string) {
	for _, pluginName := range plugins {
		switch pluginName {
		case identity.PluginName:
			p.EnablePlugin(identity.NewPlugin())
		default:
			panic(fmt.Sprintf("trying to load unknown plugin: %s", pluginName))
		}
	}
}

func loadPluginsConfiguration(p persona.Persona, configFile []byte) error {
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
