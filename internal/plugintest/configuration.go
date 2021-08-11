package plugintest

import (
	"testing"

	"github.com/endorama/devid/internal/plugin"
)

// GetTestConfig return a valid and complete plugin.Config for testing purposes.
func GetConfig(t *testing.T) plugin.Config {
	t.Helper()

	p := GetPersona(t)

	return p.Config
}
