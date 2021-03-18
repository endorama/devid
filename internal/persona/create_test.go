package persona_test

import (
	"testing"

	"github.com/endorama/devid/internal/persona"
)

func TestCreate(t *testing.T) {
	setupTestEnv()

	p, _ := persona.New("carol")

	if err := persona.Create(p); err != nil {
		t.Errorf("Unexpected error %w", err)
	}

	_ = persona.Delete(p)
}

func TestCreate_DoNotOverride(t *testing.T) {
	setupTestEnv()

	p, _ := persona.New("bob")

	if err := persona.Create(p); err == nil {
		t.Errorf("Should have been an error error but it's %w", err)
	}
}
