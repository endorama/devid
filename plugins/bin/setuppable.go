package bin

import (
	"os"
	"path"
)

func (p Plugin) Setup(profileLocation string) error {
	loc := path.Join(profileLocation, pluginName)

	return p.fs.MkdirAll(loc, 0750)
}
