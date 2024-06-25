package archive_test

import (
	"archive/tar"
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/endorama/devid/internal/archive"
)

// FailWriter is a Writer that always returns an error on writes.
type FailWriter struct{ io.Writer }

// Write implements io.Writer.
func (FailWriter) Write(_ []byte) (int, error) {
	return 0, errors.New("failed")
}

func TestArchive_Create(t *testing.T) {
	t.Run("with failing writer", func(t *testing.T) {
		err := archive.Create(&FailWriter{}, []string{"testdata/file.yaml"})
		assert.Error(t, err)
	})

	t.Run("with unexisting file", func(t *testing.T) {
		var out bytes.Buffer
		err := archive.Create(&out, []string{"testdata/doesnotexists.yaml"})
		assert.Error(t, err)
	})

	t.Run("successful", func(t *testing.T) {
		var out bytes.Buffer
		err := archive.Create(&out, []string{"testdata/file.yaml"})
		assert.NoError(t, err, "no error expected here")

		// test reading from the resulting buffer
		// this test does not check the content, only that a valid tar has been created
		tarReader := tar.NewReader(&out)
		_, err = tarReader.Next()
		assert.NoError(t, err, "reading should work")
	})
}
