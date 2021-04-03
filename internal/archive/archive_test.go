package archive_test

import (
	"archive/tar"
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/endorama/devid/internal/archive"
	"github.com/endorama/devid/internal/utils"
)

func TestArchive_Create(t *testing.T) {
	files, err := utils.Walk("testdata")
	if err != nil {
		panic(err)
	}

	assert.ElementsMatch(t, files, []string{"testdata/file.yaml"}, "do not match")

	var out bytes.Buffer
	err = archive.Create(&out, files)
	assert.NoError(t, err, "no error expected here")

	// test reading from the resulting buffer
	// this test does not check the content, only that a valid tar has been created
	tarReader := tar.NewReader(&out)
	_, err = tarReader.Next()
	assert.NoError(t, err, "reading should work")
}
