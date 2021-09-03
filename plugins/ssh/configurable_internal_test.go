package ssh

import (
	"testing"

	"github.com/endorama/devid/internal/plugintest"
	"github.com/stretchr/testify/assert"
)

func TestPlugin_ConfigureDefaults(t *testing.T) {
	cfg := plugintest.GetConfig(t, "alice").Sub("ssh")
	assert.NotNil(t, cfg)

	p := NewPlugin()

	err := p.Configure(cfg)
	if err != nil {
		t.Errorf("cannot load plugin config: %w", err)
	}

	assert.Nil(t, err, "err should be nil")

	assert.Equal(t, p.config.CachePath, defaultCachePath)
	assert.Equal(t, p.config.Keys, defaultSshKeys)
}
