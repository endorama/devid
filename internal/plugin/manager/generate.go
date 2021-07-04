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

func generateExecutableFile(baseDir string, file plugin.GeneratedFile) error {
	generateFile(baseDir, file)
	fp := path.Join(baseDir, file.Name)
	os.Chmod(fp, 0700)

	return nil
}

func generateFile(baseDir string, file plugin.GeneratedFile) error {
	// TODO: resolve path to absolute to avoid directory traversal

	fp := path.Join(baseDir, file.Name)
	utils.PersistFile(fp, file.Content)
	os.Chmod(fp, 0600)

	return nil
}

func deleteFile(filepath string) error {
	e := os.Remove(filepath)
	if e != nil {
		return e
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
			for _, file := range genFiles.Executables {
				err = generateExecutableFile(path.Join(p.Location(), "bin"), file)
				if err != nil {
					return fmt.Errorf("plugin file generation failed: %w", err)
				}
			}

			for _, file := range genFiles.Files {
				err = generateFile(p.Location(), file)
				if err != nil {
					return fmt.Errorf("plugin file generation failed: %w", err)
				}
			}
		} else {
			for _, file := range genFiles.Executables {
				err = deleteFile(path.Join(p.Location(), "bin", file.Name))
				if err != nil {
					return fmt.Errorf("plugin file generation failed: %w", err)
				}
			}

			for _, file := range genFiles.Files {
				err = deleteFile(path.Join(p.Location(), file.Name))
				if err != nil {
					return fmt.Errorf("plugin file generation failed: %w", err)
				}
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
