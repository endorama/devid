package ssh

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/endorama/devid/internal/plugintest"
)

func TestPlugin_ConfigureDefaults(t *testing.T) {
	cfg := plugintest.GetConfig(t, "alice").Sub("ssh")
	assert.NotNil(t, cfg)

	p := NewPlugin()

	err := p.Configure(cfg)
	if err != nil {
		t.Errorf("cannot load plugin config: %v", err) //nolint:govet // testing
	}

	assert.Nil(t, err, "err should be nil")

	assert.Equal(t, p.config.CachePath, defaultCachePath)
	assert.Equal(t, p.config.Keys, defaultSSHKeys())
}
