package backup

import (
	"fmt"
	"os"

	encryptedarchive "github.com/endorama/devid/internal/archive/compressed/encrypted"
)

// Perform create an encrypted backup archive from the specified set of files, using the
// specified passphrase. To support relative paths, it changes cwd before creating the file.
// Archive will be created in the current folder.
// TODO: consider removing and use Task.Perform()
func Perform(b Task, passphrase string) (err error) {
	// NOTE: change folder to source location, as b.Files() return relative
	// paths for files to be added to the archive
	return wrapChdirRestore(b.Source, func() error {
		files, err := b.Files()
		if err != nil {
			return fmt.Errorf("failed retrieving files to backup: %w", err)
		}

		err = encryptedarchive.Create(b.Destination, files, passphrase)
		if err != nil {
			return fmt.Errorf("failed creating encrypted backup archive: %w", err)
		}

		return nil
	})
}

// wrapChdirRestore change folder to dir, execute the specified function
// and restore the working directory as set before the change.
func wrapChdirRestore(target string, f func() error) (err error) {
	var cwd string

	cwd, err = os.Getwd()
	if err != nil {
		return fmt.Errorf("cannot get current working directory: %w", err)
	}

	defer func() {
		// NOTE: In case changing back to the previous working directory
		// fails is possible to shadow the error of the passed in function.
		// This is acceptable as it is not considered common to not be able
		// to chdir to a previously working directory.
		err = os.Chdir(cwd)
	}()

	err = os.Chdir(target)
	if err != nil {
		return fmt.Errorf("cannot change dir: %w", err)
	}

	err = f()

	return err
}
