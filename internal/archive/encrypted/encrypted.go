package encrypted

import (
	"fmt"
	"io"

	"filippo.io/age"

	compressedarchive "github.com/endorama/devid/internal/archive/compressed"
)

// Create creates an encrypted compressed tar archive
func Create(out io.Writer, files []string, password string) error {
	r, err := age.NewScryptRecipient(password)
	if err != nil {
		return fmt.Errorf("failed creating scrypt recipient: %w", err)
	}

	w, err := age.Encrypt(out, r)
	if err != nil {
		return fmt.Errorf("failed encrypting: %w", err)
	}
	defer w.Close()

	err = compressedarchive.Create(w, files)
	if err != nil {
		return fmt.Errorf("cannot create archive: %w", err)
	}

	return nil
}
