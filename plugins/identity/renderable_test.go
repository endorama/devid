package identity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/endorama/devid/internal/plugintest"
	"github.com/endorama/devid/plugins/identity"
)

func TestPlugin_Render(t *testing.T) {
	p := plugintest.GetPersona(t, "alice")

	i := identity.NewPlugin()

	icfg := p.Config.Sub(identity.PluginName)
	assert.NotNil(t, icfg)

	err := i.Configure(icfg)
	assert.Nil(t, err)

	r := i.Render(p.Name(), p.Location())

	expected := `export IDENTITY_EMAIL="alice@example.com"
export IDENTITY_NAME="Alice Testcase"
`
	assert.Equal(t, expected, r)
}
