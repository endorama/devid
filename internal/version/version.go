package version

import (
	"fmt"
)

var (
	version = "master"  //nolint:gochecknoglobals
	commit  = "unknown" //nolint:gochecknoglobals
)

// BuildString returns full version information.
func BuildString() string {
	return fmt.Sprintf("%s (from commit %s)", version, commit)
}

// Commit return current commit value.
func Commit() string {
	return commit
}

// Version return current version value.
func Version() string {
	return version
}
