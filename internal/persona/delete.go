package persona

import (
	"fmt"
	"os"
)

// Delete deletes specified persona configuration
// It does not delete files within the persona's folder.
func Delete(p Persona) error {
	if !p.Exists() {
		return nil
	}

	if err := os.Remove(p.File()); err != nil {
		return fmt.Errorf("cannot delete persona(%s): %w", p.Name(), err)
	}

	return nil
}
