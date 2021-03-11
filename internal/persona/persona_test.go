package persona_test

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/endorama/devid/internal/persona"
)

func setupTestEnv() {
	cwd, _ := os.Getwd()
	personasLocation := path.Join(cwd, "..", "..", "test", "testdata")
	os.Setenv("DEVID_PERSONAS_LOCATION", personasLocation)
}

func TestPersona_DoNotExistsWithoutFolder(t *testing.T) {
	setupTestEnv()

	p, _ := persona.New("donotexists", os.Getenv("DEVID_PERSONAS_LOCATION"))

	fmt.Println(p)
	if p.Exists() {
		t.Errorf("Persona.Exists() = %v, want %v", p.Exists(), false)
	}
}

func TestPersona_DoNotExistsWithFolder(t *testing.T) {
	setupTestEnv()

	p, _ := persona.New("alice", os.Getenv("DEVID_PERSONAS_LOCATION"))

	fmt.Println(p)
	if p.Exists() {
		t.Errorf("Persona.Exists() = %v, want %v", p.Exists(), false)
	}
}
