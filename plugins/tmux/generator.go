package tmux

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/endorama/devid/internal/plugin"
)

func (p *Plugin) Generate(_ string) (plugin.Generated, error) {
	tmux := strings.Builder{}
	tmux.WriteString("#!/usr/bin/env bash\n")
	tmux.WriteString("exec ")

	// we need to use the absolute path or will end up in a loop
	systemTmuxPath, err := exec.LookPath("tmux")
	if err != nil {
		return plugin.Generated{}, fmt.Errorf("cannot lookup tmux path: %w", err)
	}

	tmux.WriteString(systemTmuxPath)

	tmux.WriteString(" -S \"/tmp/devid-$TMUX_SOCKET_NAME\" ")
	tmux.WriteString("\"$@\"")

	// binFilePath := path.Join(personaDirectory, "bin", "tmux")
	// utils.PersistFile(binFilePath, tmux.String())
	// os.Chmod(binFilePath, 0700)

	return plugin.Generated{
		Executables: []plugin.GeneratedFile{
			{Name: "tmux", Content: tmux.String()},
		}}, nil
}
