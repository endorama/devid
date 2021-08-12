package envs_test

import (
	"testing"

	"github.com/endorama/devid/internal/plugin"
	"github.com/endorama/devid/internal/plugintest"
	"github.com/endorama/devid/plugins/envs"
	"github.com/stretchr/testify/assert"
)

func TestPlugin_Configurable(t *testing.T) {
	p := envs.NewPlugin()
	assert.Implements(t, (*plugin.Configurable)(nil), p)
}

func TestPlugin_Config(t *testing.T) {
	cfg := plugintest.GetConfig(t)

	p := envs.NewPlugin()

	if err := p.LoadConfig(cfg); err != nil {
		t.Errorf("cannot load plugin config: %w", err)
	}

	assert.ObjectsAreEqual(cfg, p.Config().(envs.Config))
}

func TestPlugin_LoadConfig(t *testing.T) {
	cfg := plugintest.GetConfig(t)

	p := envs.NewPlugin()

	err := p.LoadConfig(cfg)
	if err != nil {
		t.Errorf("cannot load plugin config: %w", err)
	}

	assert.Nil(t, err, "err should be nil")
}
