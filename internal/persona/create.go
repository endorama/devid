package persona

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var errPersonaExists = errors.New("persona already exists")

// Create creates specific persona configuration in the personas_location folder
// It does not override an existing persona
// If a folder with the same name exists but is not a persona, proceeds.
func Create(p Persona) error {
	if p.Exists() {
		return errPersonaExists
	}

	if _, err := os.Stat(p.Location()); os.IsNotExist(err) {
		log.Printf("%s does not exists, creating\n", p.Location())

		if err := os.Mkdir(p.Location(), 0755); err != nil {
			return fmt.Errorf("cannot create directory: %w", err)
		}
	}

	err := p.Config.SafeWriteConfig()
	if err != nil {
		return fmt.Errorf("cannot write persona's config file: %w", err)
	}

	return nil
}
