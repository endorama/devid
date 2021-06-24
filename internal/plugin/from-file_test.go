package plugin_test

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"testing"

	"github.com/endorama/devid/internal/plugin"
)

func TestLoadConfigFromFile(t *testing.T) {
	log.SetOutput(ioutil.Discard)

	cwd, _ := os.Getwd()
	personasLocation := path.Join(cwd, "..", "..", "test", "testdata", "personas")

	_, err := plugin.LoadConfigFromFile(path.Join(personasLocation, "bob", "config.yaml"))

	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
}

func TestLoadConfigFromFile_missingFile(t *testing.T) {
	log.SetOutput(ioutil.Discard)

	cwd, _ := os.Getwd()
	personasLocation := path.Join(cwd, "..", "..", "test", "testdata", "personas")

	_, err := plugin.LoadConfigFromFile(path.Join(personasLocation, "alice", "config.yaml"))

	if err == nil {
		t.Error("should have returned an error, not nil")
	}
}

func TestLoadConfigFromFile_wrongFile(t *testing.T) {
	log.SetOutput(ioutil.Discard)

	cwd, _ := os.Getwd()

	_, err := plugin.LoadConfigFromFile(path.Join(cwd, "..", "..", "test", "testdata", "not-a-yaml-file.yaml"))

	if err == nil {
		t.Error("should have returned an error, not nil")
	}
}

func TestLoadConfig(t *testing.T) {
	log.SetOutput(ioutil.Discard)

	external := plugin.NewConfig()

	_, err := plugin.LoadConfig(external)

	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
}

func TestLoadConfig_wrongVersion(t *testing.T) {
	log.SetOutput(ioutil.Discard)

	external := plugin.Config{
		APIVersion: "wrong",
	}

	cfg, err := plugin.LoadConfig(external)

	if err == nil {
		t.Errorf("should have returned an error %v", cfg)
	}
}
