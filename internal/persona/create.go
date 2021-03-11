package persona

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"gopkg.in/yaml.v1"
)

func Create(p Persona) error {
	if p.Exists() {
		return fmt.Errorf("Persona already exists")
	}

	if _, err := os.Stat(p.Location()); os.IsNotExist(err) {
		log.Printf("%s does not exists, creating\n", p.Location())
		if err := os.Mkdir(p.Location(), 0755); err != nil {
			return err
		}
	}

	d, err := yaml.Marshal(&p)
	if err != nil {
		return fmt.Errorf("cannot marshal yaml: %w", err)
	}

	err = ioutil.WriteFile(path.Join(p.Location(), filename), d, 0644)
	if err != nil {
		return fmt.Errorf("cannot write file: %w", err)
	}

	return nil
}
