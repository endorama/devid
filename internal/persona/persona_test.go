package persona_test

import (
	"os"
	"path"
	"testing"

	"github.com/endorama/devid/internal/persona"
	"github.com/endorama/devid/internal/settings"
	"github.com/spf13/viper"
)

func setupTestEnv() {
	cwd, _ := os.Getwd()
	personasLocation := path.Join(cwd, "..", "..", "test", "testdata", "profiles")
	settings.Init()
	viper.Set("personas_location", personasLocation)
}

func TestPersona(t *testing.T) {

}

func TestPersona_DoNotExistsWithoutFolder(t *testing.T) {
	setupTestEnv()

	p, _ := persona.New("donotexists")

	if p.Exists() {
		t.Errorf("Persona.Exists() = %v, want %v", p.Exists(), false)
	}
}

func TestPersona_DoNotExistsWithFolder(t *testing.T) {
	setupTestEnv()

	p, _ := persona.New("alice")

	if p.Exists() {
		t.Errorf("Persona.Exists() = %v, want %v", p.Exists(), false)
	}
}

func TestPersona_DoExists(t *testing.T) {
	setupTestEnv()

	p, _ := persona.New("bob")

	if !p.Exists() {
		t.Errorf("Persona.Exists() = %v, want %v", p.Exists(), true)
	}
}
