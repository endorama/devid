package plugintest

import (
	"io/ioutil"
	"log"
	"path"
	"testing"

	"github.com/spf13/viper"

	"github.com/endorama/devid/internal/persona"
)

func GetPersona(t *testing.T, name string) persona.Persona {
	t.Helper()

	if !testing.Verbose() {
		log.SetOutput(ioutil.Discard)
	}

	// NOTE: this loads personas from within a testdata folder within the tested
	// package. Each package should supply it's own personas for testing
	personasLocation := path.Join("testdata")
	viper.Set("personas_location", personasLocation)

	p, err := persona.Load(name)
	if err != nil {
		t.Fatalf("cannot load test persona: %s", err)
	}

	return p
}
