package identity_test

import (
	"testing"

	"github.com/endorama/devid/internal/plugintest"
	"github.com/endorama/devid/plugins/identity"
	"github.com/stretchr/testify/assert"
)

func TestPlugin_Configure(t *testing.T) {
	config := plugintest.GetConfig(t, "alice").Sub("identity")
	assert.NotNil(t, config)

	p := identity.NewPlugin()

	err := p.Configure(config)
	if err != nil {
		t.Errorf("cannot load plugin config: %w", err)
	}

	assert.Nil(t, err, "err should be nil")
}
