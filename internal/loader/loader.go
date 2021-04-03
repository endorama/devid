package loader

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/endorama/devid/internal/persona"
	"github.com/endorama/devid/plugins/identity"
	"gopkg.in/yaml.v1"
)

var errUnsupportedAPIVersion = errors.New("unsupported API version")
var errUnknownPlugin = errors.New("trying to load unknown plugin")

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
		return errUnsupportedAPIVersion
	}

	p.APIVersion = fromConfig.APIVersion

	for _, pluginName := range fromConfig.Plugins {
		switch pluginName {
		case identity.PluginName:
			p.EnablePlugin(identity.NewPlugin())
		default:
			return fmt.Errorf("%w: %s", errUnknownPlugin, pluginName)
		}
	}

	err = loadPluginsConfiguration(*p, yamlFile)
	if err != nil {
		return err
	}

	return nil
}
