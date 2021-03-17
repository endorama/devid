package archive

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
)

// Create creates an encrypted gzipped tar archive file.
func Create(out io.Writer, files []string) error {
	tw := tar.NewWriter(out)
	defer tw.Close()

	errs := []error{}
	for _, file := range files {
		err := addToArchive(tw, file)
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) != 0 {
		return fmt.Errorf("cannot create archive: %s", errs)
	}

	return nil
}

func addToArchive(tw *tar.Writer, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("cannot open file: %w", err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("cannot stat file: %w", err)
	}

	// Create a tar Header from the FileInfo data
	header, err := tar.FileInfoHeader(info, info.Name())
	if err != nil {
		return fmt.Errorf("cannot create tar file header: %w", err)
	}

	// NOTE: Use full path as name (FileInfoHeader only takes the basename)
	// If we don't do this the directory structure would
	// not be preserved
	// https://golang.org/src/archive/tar/common.go?#L626
	header.Name = filename

	err = tw.WriteHeader(header)
	if err != nil {
		return fmt.Errorf("cannot write tar header: %w", err)
	}

	_, err = io.Copy(tw, file)
	if err != nil {
		return fmt.Errorf("cannot copy file to tar: %w", err)
	}

	return nil
}
