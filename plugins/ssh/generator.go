package ssh

import (
	"fmt"
	"path"
	"strings"

	"github.com/endorama/devid/internal/plugin"
)

// Generate a file in the bin plugin directory.
// Implements the Generator interface.
func (p *Plugin) Generate(personaDirectory string) (plugin.Generated, error) {
	sshPluginFolder := path.Join(personaDirectory, pluginName)

	knownHostsFile := path.Join(sshPluginFolder, "known_hosts")
	knownHostsOption := fmt.Sprintf("-o UserKnownHostsFile=%s ", knownHostsFile)

	configFile := path.Join(sshPluginFolder, "config")
	configOption := fmt.Sprintf("-F %s ", configFile)

	ssh := strings.Builder{}
	ssh.WriteString("#!/usr/bin/env bash\n")
	ssh.WriteString("exec /usr/bin/ssh ")
	ssh.WriteString(knownHostsOption)
	ssh.WriteString(configOption)
	ssh.WriteString("$@")

	scp := strings.Builder{}
	scp.WriteString("#!/usr/bin/env bash\n")
	scp.WriteString("exec /usr/bin/scp ")
	scp.WriteString(knownHostsOption)
	scp.WriteString(configOption)
	scp.WriteString("$@")

	return plugin.Generated{
		Executables: []plugin.GeneratedFile{
			{Name: "ssh", Content: ssh.String()},
			{Name: "scp", Content: scp.String()},
		},
		Files: []plugin.GeneratedFile{},
	}, nil
}
