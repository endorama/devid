package compressed_test

import (
	"errors"
	"io"
	"os"
	"testing"

	"github.com/endorama/devid/internal/archive/compressed"
	"github.com/stretchr/testify/assert"
)

// FailWriter is a Writer that always returns an error on writes.
type FailWriter struct{ io.Writer }

// Write implements io.Writer.
func (_ FailWriter) Write(_ []byte) (int, error) {
	return 0, errors.New("failed")
}

func TestCreate(t *testing.T) {
	getTempFile := func(t *testing.T) io.Writer {
		t.Helper()

		tmpFile, err := os.CreateTemp("", "devid-test-*")
		assert.NoError(t, err)

		t.Cleanup(func() {
			os.Remove(tmpFile.Name())
			tmpFile.Close()
		})

		return tmpFile
	}

	f := []string{"testdata/foo.txt"}

	tests := []struct {
		name string

		getWriter func(t *testing.T) io.Writer
		files     []string

		expectedErr bool
	}{
		{
			name:        "with writer failure",
			getWriter:   func(t *testing.T) io.Writer { return FailWriter{} },
			files:       f,
			expectedErr: true,
		},
		{
			name:        "with unexisting files",
			getWriter:   getTempFile,
			files:       []string{"testdata/doesnotexists"},
			expectedErr: true,
		},
		{
			// NOTE: does not test archive content
			name:      "successful",
			getWriter: getTempFile,
			files:     f,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := tt.getWriter(t)
			err := compressed.Create(w, tt.files)
			if tt.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
