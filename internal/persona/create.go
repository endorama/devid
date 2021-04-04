package persona

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"gopkg.in/yaml.v1"
)

var errPersonaExists = errors.New("Persona already exists")

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

	d, err := yaml.Marshal(&p)
	if err != nil {
		return fmt.Errorf("cannot marshal yaml: %w", err)
	}

	err = ioutil.WriteFile(path.Join(p.Location(), filename), d, 0600)
	if err != nil {
		return fmt.Errorf("cannot write file: %w", err)
	}

	return nil
}
