package bin

import (
	"fmt"
	"os"
	"path"
)

const binFolderPerm = os.FileMode(0750)

func (p Plugin) Setup(profileLocation string) error {
	loc := path.Join(profileLocation, pluginName)

	if err := p.fs.MkdirAll(loc, binFolderPerm); err != nil {
		return fmt.Errorf("cannot create directory tree: %v", err)
	}

	return nil
}
