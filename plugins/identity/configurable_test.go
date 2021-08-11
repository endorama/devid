package identity_test

import (
	"testing"

	"github.com/endorama/devid/internal/plugintest"
	"github.com/endorama/devid/plugins/identity"
	"github.com/stretchr/testify/assert"
)

func TestPlugin_Config(t *testing.T) {
	cfg := plugintest.GetConfig(t)

	p := identity.NewPlugin()

	if err := p.LoadConfig(cfg); err != nil {
		t.Errorf("cannot load plugin config: %w", err)
	}

	assert.ObjectsAreEqual(cfg, p.Config().(identity.Config))
}

func TestPlugin_LoadConfig(t *testing.T) {
	cfg := plugintest.GetConfig(t)

	p := identity.NewPlugin()

	err := p.LoadConfig(cfg)
	if err != nil {
		t.Errorf("cannot load plugin config: %w", err)
	}

	assert.Nil(t, err, "err should be nil")
}
