package utils

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var errEditorValueNotAllowed = errors.New("$EDITOR environment variable value is not allowed")

// OpenWithEditor open the specified path with the EDITOR set in the corresponding environment
// variable.
func OpenWithEditor(path string) error {
	editorCmd := os.ExpandEnv("$EDITOR")
	if !isEditorAllowed(editorCmd) {
		if editorCmd == "" {
			return fmt.Errorf("%w, is empty", errEditorValueNotAllowed)
		}

		return fmt.Errorf("%w (%s) allowed editors: %s",
			errEditorValueNotAllowed, editorCmd, strings.Join(AllowedEditors, ","))
	}

	editor := exec.Command(editorCmd, path)
	// NOTE: pass current STDIN and STDOUT to EDITOR so it can properly start
	editor.Stdin = os.Stdin
	editor.Stdout = os.Stdout

	if err := editor.Run(); err != nil {
		return fmt.Errorf("cannot run '%s': %w", editor.String(), err)
	}

	return nil
}

// AllowedEditors is a list of allowed values for the EDITOR environment variable.
// FIXME: prevent unknown command execution when some of these editor is not available.
// FIXME: make this a function so is not directly modifiable.
var AllowedEditors = []string{ //nolint:gochecknoglobals // implementation detail
	"/bin/ed",
	"/bin/nano",
	"/usr/bin/vim",
	"nano",
	"vim",
}

// isEditorAllowed solve gosec:G204(Audit use of command execution) by sanitizing the values of
// $EDITOR environment value to prevent arbitrary command injection.
func isEditorAllowed(editorCmd string) bool {
	allowedEditorCmds := AllowedEditors

	return containsString(allowedEditorCmds, editorCmd)
}

// containsString check if a specified string is present in a slice.
func containsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}
