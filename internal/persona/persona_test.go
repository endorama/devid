package persona_test

import (
	"os"
	"path"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	"github.com/endorama/devid/internal/persona"
	"github.com/endorama/devid/internal/settings"
)

func setupTestEnv() {
	cwd, _ := os.Getwd()
	personasLocation := path.Join(cwd, "..", "..", "test", "testdata", "personas")

	settings.Init()
	viper.Set("personas_location", personasLocation)
}

func TestPersona_File(t *testing.T) {
	setupTestEnv()

	p, _ := persona.New("dan")

	assert.Equal(t,
		path.Join(viper.GetString("personas_location"), p.Name(), "config.yaml"),
		p.File(), "they should be equal")
}

func TestPersona_Location(t *testing.T) {
	setupTestEnv()

	p, _ := persona.New("dan")

	assert.Equal(t,
		path.Join(viper.GetString("personas_location"), p.Name()),
		p.Location(), "they should be equal")
}

func TestPersona_Name(t *testing.T) {
	setupTestEnv()

	p, _ := persona.New("charlie")

	assert.Equal(t,
		"charlie",
		p.Name(), "they should be equal")
}

func TestPersona_Whoami(t *testing.T) {
	setupTestEnv()

	p, _ := persona.New("charlie")

	assert.Equal(t, "charlie", p.Whoami(), "they should be equal")
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
