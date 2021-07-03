package _template

import (
	"os"
	"path"
	"strings"

	"github.com/endorama/devid/internal/utils"
)

// Generate a file in the bin plugin directory.
// Implements the Generator interface.
func (p *Plugin) Generate(personaDirectory string) error {
	wrappedBin := strings.Builder{}
	wrappedBin.WriteString("#!/usr/bin/env bash\n")
	wrappedBin.WriteString("exec ")

	wrappedBin.WriteString("ls $@")

	wrappedBinFilePath := path.Join(personaDirectory, "bin", "wrappedBin")
	utils.PersistFile(wrappedBinFilePath, wrappedBin.String())
	os.Chmod(wrappedBinFilePath, 0700)

	return nil
}
