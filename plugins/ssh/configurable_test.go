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

func TestPlugin_LoadConfig(t *testing.T) {
	cfg := plugintest.GetConfig(t, "alice").Sub("ssh")
	assert.NotNil(t, cfg)

	p := ssh.NewPlugin()

	err := p.Configure(cfg)
	if err != nil {
		t.Errorf("cannot load plugin config: %v", err) //nolint:govet // don't care
	}

	assert.Nil(t, err, "err should be nil")
}
