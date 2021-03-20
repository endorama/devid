package backup

import (
	"io"

	"github.com/endorama/devid/internal/utils"
)

// Task contains required information for backup operation.
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
	return utils.WalkRelative(b.Source)
}
