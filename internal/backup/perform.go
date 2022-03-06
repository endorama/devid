package backup

import (
	"fmt"
	"os"

	encryptedarchive "github.com/endorama/devid/internal/archive/compressed/encrypted"
)

// Perform create an encrypted backup archive from the specified set of files, using the
// specified passphrase. To support relative paths, it allows a cwd parameter to change directory
// before creating the file.
// Archive will be created in the current folder.
func Perform(b Task, passphrase string) error {
	// NOTE: change folder to source location, as b.Files() return relative
	// paths for files to be added to the archive
	err := os.Chdir(b.Source)
	if err != nil {
		return fmt.Errorf("cannot change dir: %w", err)
	}

	files, err := b.Files()
	if err != nil {
		return fmt.Errorf("failed retrieving files to backup: %w", err)
	}

	err = encryptedarchive.Create(b.Destination, files, passphrase)
	if err != nil {
		return fmt.Errorf("failed creating encrypted backup archive: %w", err)
	}

	return nil
}
