package ssh_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/endorama/devid/internal/plugin"
	"github.com/endorama/devid/internal/plugintest"
	"github.com/endorama/devid/plugins/ssh"
)

func TestPlugin_Configurable(t *testing.T) {
	p := ssh.NewPlugin()
	assert.Implements(t, (*plugin.Configurable)(nil), p)
}

func TestPlugin_Config(t *testing.T) {
	cfg := plugintest.GetConfig(t)

	p := ssh.NewPlugin()

	if err := p.LoadConfig(cfg); err != nil {
		t.Errorf("cannot load plugin config: %w", err)
	}

	assert.ObjectsAreEqual(cfg, p.Config().(ssh.Config))
}

func TestPlugin_LoadConfig(t *testing.T) {
	cfg := plugintest.GetConfig(t)

	p := ssh.NewPlugin()

	err := p.LoadConfig(cfg)
	if err != nil {
		t.Errorf("cannot load plugin config: %w", err)
	}

	assert.Nil(t, err, "err should be nil")
}
