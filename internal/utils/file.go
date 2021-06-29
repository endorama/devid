package utils

import (
	"fmt"
	"io/ioutil"
	"os"
)

// PersistFile writes content to a path.
func PersistFile(path, content string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("cannot craete path: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("cannot write file: %w", err)
	}

	return nil
}

// PersistExecutableFile writes a content to a path and make it executable.
func PersistExecutableFile(path, content string) error {
	err := PersistFile(path, content)
	if err != nil {
		return err
	}

	err = os.Chmod(path, 0700)
	if err != nil {
		return fmt.Errorf("cannot change file permissions: %w", err)
	}

	return nil
}

func ReadFile(path string) ([]byte, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return []byte{}, fmt.Errorf("cannot read file at %s: %w", path, err)
	}

	return content, nil
}
