package persona_test

import (
	"testing"

	"github.com/endorama/devid/internal/persona"
)

func TestCreate(t *testing.T) {
	setupTestEnv()

	p, _ := persona.New("carol")

	if err := persona.Create(p); err != nil {
		t.Errorf("unexpected error %v", err)
	}

	_ = persona.Delete(p)
}

func TestCreate_DoNotOverride(t *testing.T) {
	setupTestEnv()

	p, _ := persona.New("bob")

	if err := persona.Create(p); err == nil {
		t.Errorf("should have been an error error but it's %v", err)
	}
}
