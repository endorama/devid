package name_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/endorama/devid/internal/plugin"
	"github.com/endorama/devid/internal/plugintest"
	"github.com/endorama/devid/plugins/name"
)

func TestPlugin_Configurable(t *testing.T) {
	p := name.NewPlugin()
	assert.Implements(t, (*plugin.Configurable)(nil), p)
}

func TestPlugin_Configure(t *testing.T) {
	config := plugintest.GetConfig(t, "alice").Sub("identity")
	assert.NotNil(t, config)

	p := name.NewPlugin()

	err := p.Configure(config)
	if err != nil {
		t.Errorf("cannot load plugin config: %w", err)
	}

	assert.Nil(t, err, "err should be nil")
}
