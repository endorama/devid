package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// WalkRelative walks the specified source path and return a slice of
// relative paths contained.
func WalkRelative(source string) ([]string, error) {
	files := []string{}

	absoluteFiles, err := Walk(source)
	if err != nil {
		return files, fmt.Errorf("cannot walk source: %w", err)
	}

	for _, f := range absoluteFiles {
		files = append(files,
			strings.ReplaceAll(f, fmt.Sprintf("%s/", source), ""))
	}

	return files, nil
}

var errSourceIsNonAFolder = errors.New("source is not a folder")

// Walk walks the specified source path and return a slice of absolute paths
// contained.
func Walk(source string) ([]string, error) {
	files := []string{}

	info, err := os.Stat(source)
	if os.IsNotExist(err) {
		return files, fmt.Errorf("folder does not exists: %w", err)
	}

	if !info.IsDir() {
		return files, errSourceIsNonAFolder
	}

	walkErrs := []error{}
	err = filepath.Walk(source,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				walkErrs = append(walkErrs, err)
			}

			if !info.IsDir() {
				files = append(files, path)
			}

			return nil
		})

	if err != nil {
		return files, fmt.Errorf("backup source walk failed: %w", err)
	}

	if len(walkErrs) != 0 {
		//nolint:goerr113 // this error includes all errors from the walk operation
		return files, fmt.Errorf("errors walking source: %s", walkErrs)
	}

	return files, nil
}
