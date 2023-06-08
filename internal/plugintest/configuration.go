package plugintest

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

// GetTestConfig return a valid and complete plugin.Config for testing purposes.
func GetConfig(t *testing.T, name string) *viper.Viper {
	t.Helper()

	p := GetPersona(t, name)

	return p.Config
}

// IsEnabled checks if the specific plugin is enabled in the provided config.
func IsEnabled(t *testing.T, plugin string, config *viper.Viper) bool {
	t.Helper()
	return config.GetBool(fmt.Sprintf("%s.enabled", plugin))
}
