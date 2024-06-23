package backup

import (
	"fmt"
	"io"

	"github.com/endorama/devid/internal/utils"
)

// Task contains required information for backup operation.
// TODO: add Perform() method to run backups
// TODO: move New() here
type Task struct {
	// Name is the name of the backup task
	Name string
	// Source is the source location to backup
	Source string
	// Destination is the location to write the backup to
	Destination io.Writer
}

// Files retrieves list of files to backup by walking Source.
func (b Task) Files() ([]string, error) {
	paths, err := utils.WalkRelative(b.Source)
	if err != nil {
		return paths, fmt.Errorf("cannot walk relative path: %v", err)
	}

	return paths, nil
}
