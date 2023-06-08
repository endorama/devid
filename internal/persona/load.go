package persona

import (
	"errors"
	"fmt"
)

var errPersonaDoesNotExists = errors.New("does not exists")

func Load(name string) (Persona, error) {
	p, err := New(name)
	if err != nil {
		return p, fmt.Errorf("init failed: %w", err)
	}

	if !p.Exists() {
		return p, fmt.Errorf("%w in %s", errPersonaDoesNotExists, p.Location())
	}

	err = p.Load()
	if err != nil {
		return p, fmt.Errorf("cannot load persona configuration: %w", err)
	}

	return p, nil
}
