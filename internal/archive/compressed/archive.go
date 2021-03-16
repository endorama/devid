package compressed

import (
	"compress/gzip"
	"fmt"
	"io"

	"github.com/endorama/devid/internal/archive"
)

// Create creates a compressed tar archive file
func Create(out io.Writer, files []string) error {
	gw := gzip.NewWriter(out)
	defer gw.Close()

	err := archive.Create(gw, files)
	if err != nil {
		return fmt.Errorf("cannot create compressed archive: %s", err)
	}

	return nil
}
