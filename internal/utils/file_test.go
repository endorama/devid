package utils_test

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/endorama/devid/internal/utils"

	"github.com/stretchr/testify/assert"
)

const testContent = "foobar"

func TestPersistFile(t *testing.T) {
	t.Run("when file doesn't exists", func(t *testing.T) {
		p := fmt.Sprintf("testdata/%s.txt", t.Name())
		err := os.MkdirAll(path.Dir(p), os.FileMode(0770))
		assert.NoError(t, err)

		err = utils.PersistFile(p, testContent)
		assert.NoError(t, err)

		c, err := os.ReadFile(p)
		assert.NoError(t, err)
		assert.Equal(t, testContent, string(c))
		os.Remove(p)
	})

	t.Run("when the file exists", func(t *testing.T) {
		p := fmt.Sprintf("testdata/%s.txt", t.Name())
		err := os.MkdirAll(path.Dir(p), os.FileMode(0770))
		assert.NoError(t, err)
		f, err := os.Create(p)
		assert.NoError(t, err)
		f.Close()

		err = utils.PersistFile(p, testContent)
		assert.NoError(t, err)

		c, err := os.ReadFile(p)
		assert.NoError(t, err)
		assert.Equal(t, testContent, string(c))
		os.Remove(p)
	})
}

func TestPersistExecutableFile(t *testing.T) {
	p := fmt.Sprintf("testdata/%s.txt", t.Name())
	err := os.MkdirAll(path.Dir(p), os.FileMode(0770))
	assert.NoError(t, err)
	f, err := os.Create(p)
	assert.NoError(t, err)
	f.Close()

	err = utils.PersistExecutableFile(p, testContent)
	assert.NoError(t, err)

	c, err := os.ReadFile(p)
	assert.NoError(t, err)
	assert.Equal(t, testContent, string(c))

	info, err := os.Stat(p)
	assert.NoError(t, err)
	assert.Equal(t, os.FileMode(0700), info.Mode())

	os.Remove(p)
}

func TestReadFile(t *testing.T) {
	p := fmt.Sprintf("testdata/%s.txt", t.Name())
	err := os.MkdirAll(path.Dir(p), os.FileMode(0770))
	assert.NoError(t, err)
	err = os.WriteFile(p, []byte(testContent), os.FileMode(0660))
	assert.NoError(t, err)

	c, err := utils.ReadFile(p)
	assert.NoError(t, err)
	assert.Equal(t, testContent, string(c))

	os.Remove(p)
}

func TestDeleteFile(t *testing.T) {
	p := fmt.Sprintf("testdata/%s.txt", t.Name())
	err := os.MkdirAll(path.Dir(p), os.FileMode(0770))
	assert.NoError(t, err)
	err = os.WriteFile(p, []byte(testContent), os.FileMode(0660))
	assert.NoError(t, err)

	err = utils.DeleteFile(p)
	assert.NoError(t, err)

	err = os.Remove(p)
	assert.ErrorIs(t, err, os.ErrNotExist)
}
