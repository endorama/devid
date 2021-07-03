package tmux

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/endorama/devid/internal/utils"
)

func (p *Plugin) Generate(personaDirectory string) error {
	tmux := strings.Builder{}
	tmux.WriteString("#!/usr/bin/env bash\n")
	tmux.WriteString("exec ")

	// we need to use the absolute path or will end up in a loop
	systemTmuxPath, err := exec.LookPath("tmux")
	if err != nil {
		return fmt.Errorf("cannot lookup tmux path: %w", err)
	}
	tmux.WriteString(systemTmuxPath)

	tmux.WriteString(" -S \"/tmp/devid-$TMUX_SOCKET_NAME\" ")
	tmux.WriteString("\"$@\"")

	binFilePath := path.Join(personaDirectory, "bin", "tmux")
	utils.PersistFile(binFilePath, tmux.String())
	os.Chmod(binFilePath, 0700)

	return nil
}
