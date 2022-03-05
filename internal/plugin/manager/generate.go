package manager

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/endorama/devid/internal/persona"
	"github.com/endorama/devid/internal/plugin"
	"github.com/endorama/devid/internal/utils"
)

const permUserRWX = os.FileMode(0700)
const permUserRW = os.FileMode(0600)

func generateExecutableFile(baseDir string, file plugin.GeneratedFile) error {
	generateFile(baseDir, file)

	fp := path.Join(baseDir, file.Name)
	if err := utils.PersistFile(fp, file.Content); err != nil {
		return fmt.Errorf("cannot persist file %s: %w", fp, err)
	}

	if err := os.Chmod(fp, permUserRWX); err != nil {
		return fmt.Errorf("cannot change permissions to %s on %s: %v", permUserRWX, fp, err)
	}

	return nil
}

func generateFile(baseDir string, file plugin.GeneratedFile) error {
	// TODO: resolve path to absolute to avoid directory traversal
	fp := path.Join(baseDir, file.Name)
	if err := utils.PersistFile(fp, file.Content); err != nil {
		return fmt.Errorf("cannot persist file %s: %w", fp, err)
	}

	if err := os.Chmod(fp, permUserRW); err != nil {
		return fmt.Errorf("cannot change permissions to %s on %s: %v", permUserRW, fp, err)
	}

	return nil
}

func deleteFile(filepath string) error {
	if err := os.Remove(filepath); err != nil {
		return fmt.Errorf("cannot delete file %s: %w", filepath, err)
	}

	return nil
}

func createPluginFiles(loc string, files plugin.Generated) error {
	for _, file := range files.Executables {
		err := generateExecutableFile(path.Join(loc, "bin"), file)
		if err != nil {
			return fmt.Errorf("cannot create executable file: %w", err)
		}
	}

	for _, file := range files.Files {
		err := generateFile(loc, file)
		if err != nil {
			return fmt.Errorf("cannot create file: %w", err)
		}
	}

	return nil
}

func deletePluginFiles(loc string, files plugin.Generated) error {
	for _, file := range files.Executables {
		err := deleteFile(path.Join(loc, "bin", file.Name))
		if err != nil {
			return fmt.Errorf("cannot delete executable file: %w", err)
		}
	}

	for _, file := range files.Files {
		err := deleteFile(path.Join(loc, file.Name))
		if err != nil {
			return fmt.Errorf("cannot delete file: %w", err)
		}
	}

	return nil
}

func generatePlugin(p persona.Persona, plg Plugin) error {
	if generatorPlugin, ok := plg.Instance.(plugin.Generator); ok {
		log.Printf("running generation for: %s", plg.Instance.Name())

		genFiles, err := generatorPlugin.Generate(p.Location())
		if err != nil {
			return fmt.Errorf("plugin generation failed: %w", err)
		}

		if plg.Enabled {
			err := createPluginFiles(p.Location(), genFiles)
			if err != nil {
				return fmt.Errorf("cannot create plugin generated files: %w", err)
			}
		} else {
			err := deletePluginFiles(p.Location(), genFiles)
			if err != nil {
				return fmt.Errorf("cannot delete plugin generated files: %w", err)
			}
		}
	}

	return nil
}

func Generate(p persona.Persona) ([]error, error) {
	log.SetPrefix("plugins-generate ")
	defer log.SetPrefix("")

	errs := []error{}

	for _, plg := range plugins {
		err := generatePlugin(p, plg)
		if err != nil {
			errs = append(errs, err)
		}
	}

	return errs, nil
}
