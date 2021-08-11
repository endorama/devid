package identity_test

import (
	"testing"

	"github.com/endorama/devid/internal/plugin"
	"github.com/endorama/devid/plugins/identity"
	"github.com/stretchr/testify/assert"
)

func TestNewPlugin(t *testing.T) {
	p := identity.NewPlugin()
	assert.Implements(t, (*plugin.Pluggable)(nil), p)
	assert.Implements(t, (*plugin.Configurable)(nil), p)
	assert.Implements(t, (*plugin.Renderable)(nil), p)
}

func TestPlugin_Name(t *testing.T) {
	p := identity.NewPlugin()
	assert.Equal(t, identity.PluginName, p.Name())
}

func TestPlugin_Whoami(t *testing.T) {
	t.Skip("not implemented")
}
