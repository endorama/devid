package ssh

import (
	"fmt"
	"os"
	"path"
)

const sshFolderPerm = os.FileMode(0750)

func (p Plugin) Setup(profileLocation string) error {
	loc := path.Join(profileLocation, pluginName)

	err := p.fs.MkdirAll(loc, sshFolderPerm)
	if err != nil {
		return fmt.Errorf("cannot create ssh plugin folder: %w", err)
	}

	knownHostsFile, err := os.Create(path.Join(loc, "known_hosts"))
	if err != nil {
		return fmt.Errorf("cannot create known_hosts file: %w", err)
	}

	knownHostsFile.Close()

	configFile, err := os.Create(path.Join(loc, "config"))
	if err != nil {
		return fmt.Errorf("cannot create config file: %w", err)
	}

	configFile.Close()

	return nil
}
