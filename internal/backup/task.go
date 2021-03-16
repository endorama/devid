package backup

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Task contains required information for backup operation
type Task struct {
	Name        string
	Source      string
	Destination *os.File
}

// Files retrieves list of files to backup
func (b Task) Files() ([]string, error) {
	files := []string{}
	err := filepath.Walk(b.Source,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				// strip prefix so file path is relative to source root
				files = append(files, strings.ReplaceAll(path, fmt.Sprintf("%s/", b.Source), ""))
			}
			return nil
		})
	if err != nil {
		return files, fmt.Errorf("backup source walk failed: %w", err)
	}
	return files, nil
}
