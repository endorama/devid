package persona_test

import (
	"testing"

	"github.com/endorama/devid/internal/persona"
	"github.com/spf13/viper"
)

func TestCreate(t *testing.T) {
	setupTestEnv()

	p, _ := persona.New("carol", viper.GetString("personas_location"))
	err := persona.Create(p)
	if err != nil {
		t.Errorf("Unexpected error %w", err)
	}
	_ = persona.Delete(p)
}

func TestCreate_DoNotOverride(t *testing.T) {
	setupTestEnv()

	p, _ := persona.New("bob", viper.GetString("personas_location"))
	err := persona.Create(p)
	if err == nil {
		t.Errorf("Should have been an error error but it's %w", err)
	}
}
