package plugin

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v1"
)

var errUnsupportedAPIVersion = errors.New("unsupported API version")

func LoadConfigFromFile(path string) (Config, error) {
	log.Println(fmt.Sprintf("loading config from file: %s", path))

	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("cannot read configuration file: %w", err)
	}

	log.Println(fmt.Sprintf("config: \n%+v", yamlFile))

	var fromFile Config

	err = yaml.Unmarshal(yamlFile, &fromFile)
	if err != nil {
		return Config{}, fmt.Errorf("cannot unmarshal file: %w", err)
	}

	log.Println(fmt.Sprintf("fromFile: %+v", fromFile))

	return LoadConfig(fromFile)
}

func LoadConfig(external Config) (Config, error) {
	config := NewConfig()
	if external.APIVersion != config.APIVersion {
		return config, fmt.Errorf("%w, found %s expected %s", errUnsupportedAPIVersion,
			external.APIVersion, config.APIVersion)
	}

	config.Identity = external.Identity
	config.Envs = external.Envs
	config.Tmux = external.Tmux

	log.Printf("loaded config: %+v\n", config)

	return config, nil
}
