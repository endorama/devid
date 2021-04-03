package utils

import "os"

// PersistFile writes content to a path.
func PersistFile(path, content string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString(content)
	return nil
}

// PersistExecutableFile writes a content to a path and make it executable.
func PersistExecutableFile(path, content string) error {
	err := PersistFile(path, content)
	if err != nil {
		return err
	}

	os.Chmod(path, 0700)
	return nil
}
