package persona_test

import (
	"testing"

	"github.com/endorama/devid/internal/persona"
)

func TestPersonaNew(t *testing.T) {
	p, _ := persona.New("test", "someplace")

	if p.Name() != "test" {
		t.Errorf("p.Name() = %v, want %v", p.Name(), "test")
	}
	if p.Location() != "someplace/test" {
		t.Errorf("p.Location() = %v, want %v", p.Location(), "someplace/test")
	}
}

func TestPersonaNew_Error(t *testing.T) {
	_, err := persona.New("test", "someplace")
	if err != nil {
		t.Errorf("err should be nil, got %v", err)
	}
}
