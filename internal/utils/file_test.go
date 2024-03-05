package utils_test

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/endorama/devid/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestPersistFile(t *testing.T) {
	content := "foobar"

	t.Run("when file doesn't exists", func(t *testing.T) {
		p := fmt.Sprintf("testdata/%s.txt", t.Name())
		os.MkdirAll(path.Dir(p), os.FileMode(0770))

		utils.PersistFile(p, content)

		c, err := os.ReadFile(p)
		assert.NoError(t, err)
		assert.Equal(t, content, string(c))
		os.Remove(p)
	})

	t.Run("when the file exists", func(t *testing.T) {
		p := fmt.Sprintf("testdata/%s.txt", t.Name())
		os.MkdirAll(path.Dir(p), os.FileMode(0770))
		f, err := os.Create(p)
		f.Close()

		utils.PersistFile(p, content)

		c, err := os.ReadFile(p)
		assert.NoError(t, err)
		assert.Equal(t, content, string(c))
		os.Remove(p)
	})
}

func TestPersistExecutableFile(t *testing.T) {
	content := "foobar"

	p := fmt.Sprintf("testdata/%s.txt", t.Name())
	os.MkdirAll(path.Dir(p), os.FileMode(0770))
	f, err := os.Create(p)
	f.Close()

	utils.PersistExecutableFile(p, content)

	c, err := os.ReadFile(p)
	assert.NoError(t, err)
	assert.Equal(t, content, string(c))

	info, err := os.Stat(p)
	assert.NoError(t, err)
	assert.Equal(t, os.FileMode(0700), info.Mode())

	os.Remove(p)
}

func TestReadFile(t *testing.T) {
	content := "foobar"

	p := fmt.Sprintf("testdata/%s.txt", t.Name())
	os.MkdirAll(path.Dir(p), os.FileMode(0770))
	os.WriteFile(p, []byte(content), os.FileMode(0660))

	c, err := utils.ReadFile(p)
	assert.NoError(t, err)
	assert.Equal(t, content, string(c))

	os.Remove(p)
}

func TestDeleteFile(t *testing.T) {
	content := "foobar"

	p := fmt.Sprintf("testdata/%s.txt", t.Name())
	os.MkdirAll(path.Dir(p), os.FileMode(0770))
	os.WriteFile(p, []byte(content), os.FileMode(0660))

	err := utils.DeleteFile(p)
	assert.NoError(t, err)

	err = os.Remove(p)
	assert.ErrorIs(t, err, os.ErrNotExist)
}
