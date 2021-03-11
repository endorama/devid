package utils

import (
	"os"
	"os/exec"
)

// OpenWithEditor open the specified path with the EDITOR set in the corresponding environment
// variable
func OpenWithEditor(path string) error {
	editor := exec.Command(os.ExpandEnv("$EDITOR"), path)
	// NOTE: pass current STDIN and STDOUT to EDITOR so it can properly start
	editor.Stdin = os.Stdin
	editor.Stdout = os.Stdout

	err := editor.Run()
	return err
}
