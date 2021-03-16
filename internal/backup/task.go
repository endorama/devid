package backup

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Task contains required information for backup operation
type Task struct {
	// Name is the name of the backup task
	Name string
	// Source is the source location to backup
	Source string
	// Destination is the location to write the backup to
	Destination io.Writer
}

// Files retrieves list of files to backup by walking Source
func (b Task) Files() ([]string, error) {
	files := []string{}

	info, err := os.Stat(b.Source)
	if os.IsNotExist(err) {
		return files, fmt.Errorf("folder does not exists: %w", err)
	}
	if !info.IsDir() {
		return files, fmt.Errorf("source is not a folder")
	}

	walkErrs := []error{}
	err = filepath.Walk(b.Source,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				walkErrs = append(walkErrs, err)
			}
			if !info.IsDir() {
				// strip prefix so file path is relative to source root
				files = append(files,
					strings.ReplaceAll(path, fmt.Sprintf("%s/", b.Source), ""))
			}
			return nil
		})
	if err != nil {
		return files, fmt.Errorf("backup source walk failed: %w", err)
	}
	if len(walkErrs) != 0 {
		return files, fmt.Errorf("errors walking source: %s", walkErrs)
	}

	return files, nil
}
