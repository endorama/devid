package backup

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/endorama/devid/internal/archive"
)

// Info contains required information for backup operation
type Info struct {
	Name        string
	Source      string
	Destination *os.File
}

// New initialize a BackupInfo
func New(name, source string, destination *os.File) Info {
	return Info{name, source, destination}
}

// Files retrieves list of files to backup
func (b Info) Files() ([]string, error) {
	files := []string{}
	err := filepath.Walk(b.Source,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				files = append(files, strings.ReplaceAll(path, fmt.Sprintf("%s/", b.Source), ""))
			}
			return nil
		})
	if err != nil {
		return files, fmt.Errorf("backup source walk failed: %w", err)
	}
	return files, nil
}

// Backup create an encrypted backup archive from the specified set of files, using the
// specified passphrase. To support relative paths, it allows a cwd parameter to change directory
// before creating the file.
// Archive will be created in the current folder
// func EncryptedBackup(name, password, cwd string, files []string) error {
func Backup(b Info) error {
	// create output file
	out, err := os.Create(fmt.Sprintf("%s.tar.gz", b.Name))
	if err != nil {
		return fmt.Errorf("cannot create archive: %w", err)
	}
	defer out.Close()

	// change folder to specified location, so is possible to use relative
	// paths in the archive
	err = os.Chdir(b.Source)
	if err != nil {
		return err
	}

	files, err := b.Files()
	if err != nil {
		return fmt.Errorf("failed retrieving files to backup: %w", err)
	}
	// perform archive creation, compression
	err = archive.Create(out, files)
	if err != nil {
		return fmt.Errorf("failed creating encrypted backup archive: %w", err)
	}

	return nil
}
