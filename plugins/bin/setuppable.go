package bin

import (
	"os"
	"path"
)

const binFolderPerm = os.FileMode(0750)

func (p Plugin) Setup(profileLocation string) error {
	loc := path.Join(profileLocation, pluginName)

	return p.fs.MkdirAll(loc, binFolderPerm)
}
