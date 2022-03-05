package identity

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/endorama/devid/internal/plugintest"
)

func TestPlugin_ConfigureErrorOnMissingEmail(t *testing.T) {
	cfg := plugintest.GetConfig(t, "bob").Sub("identity")
	assert.NotNil(t, cfg)

	p := NewPlugin()

	err := p.Configure(cfg)
	assert.ErrorIs(t, err, errEmailMissing)

	assert.Equal(t, p.config.Name, "Bob Testcase")
}

func TestPlugin_ConfigureErrorOnMissingName(t *testing.T) {
	cfg := plugintest.GetConfig(t, "carol").Sub("identity")
	assert.NotNil(t, cfg)

	p := NewPlugin()

	err := p.Configure(cfg)
	assert.ErrorIs(t, err, errNameMissing)

	assert.Equal(t, p.config.Email, "carol@example.com")
}
