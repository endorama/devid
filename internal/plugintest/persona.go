package plugintest

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"testing"

	"github.com/endorama/devid/internal/persona"
	"github.com/spf13/viper"
)

func GetPersona(t *testing.T) persona.Persona {
	t.Helper()

	if !testing.Verbose() {
		log.SetOutput(ioutil.Discard)
	}

	// NOTE: this is not going to work if the tests are run from a directory
	// that is not 2 level below the test/ one.
	// Note that running go test <module> change the working directory.
	cwd, _ := os.Getwd()
	personasLocation := path.Join(cwd, "..", "..", "test", "testdata", "personas")

	viper.Set("personas_location", personasLocation)

	p, err := persona.Load("bob")
	if err != nil {
		t.Fatalf("cannot load test persona: %s", err)
	}

	return p
}
