package archive

import (
	"bytes"
	"fmt"
	"io"
)

// Create creates an encrypted gzipped tar archive file
func Create(out io.Writer, files []string) error {
	var targz bytes.Buffer

	err := createArchive(files, &targz)

	if err != nil {
		return fmt.Errorf("cannot create archive: %w", err)
	}
	//
	if _, err := io.Copy(out, &targz); err != nil {
		return fmt.Errorf("failed copying data: %w", err)
	}
	// if err := w.Close(); err != nil {
	//   return fmt.Errorf("failed closing writer: %w", err)
	// }
	//
	return nil
}
