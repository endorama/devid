package plugintest

import (
	"testing"

	"github.com/spf13/viper"
)

// GetTestConfig return a valid and complete plugin.Config for testing purposes.
func GetConfig(t *testing.T, name string) *viper.Viper {
	t.Helper()

	p := GetPersona(t, name)

	return p.Config
}
