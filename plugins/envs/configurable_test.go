package envs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/endorama/devid/internal/plugin"
	"github.com/endorama/devid/internal/plugintest"
	"github.com/endorama/devid/plugins/envs"
)

func TestPlugin_Configurable(t *testing.T) {
	p := envs.NewPlugin()
	assert.Implements(t, (*plugin.Configurable)(nil), p)
}

func TestPlugin_Configure(t *testing.T) {
	cfg := plugintest.GetConfig(t, "alice").Sub("envs")
	assert.NotNil(t, cfg)

	p := envs.NewPlugin()

	err := p.Configure(cfg)
	if err != nil {
		t.Errorf("cannot load plugin config: %v", err)
	}

	assert.Nil(t, err, "err should be nil")
}
