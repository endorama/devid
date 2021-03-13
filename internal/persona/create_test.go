package persona_test

import (
	"testing"

	"github.com/endorama/devid/internal/persona"
)

func TestCreate(t *testing.T) {
	setupTestEnv()

	p, _ := persona.New("carol")
	err := persona.Create(p)
	if err != nil {
		t.Errorf("Unexpected error %w", err)
	}
	_ = persona.Delete(p)
}

func TestCreate_DoNotOverride(t *testing.T) {
	setupTestEnv()

	p, _ := persona.New("bob")
	err := persona.Create(p)
	if err == nil {
		t.Errorf("Should have been an error error but it's %w", err)
	}
}
