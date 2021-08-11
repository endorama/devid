package identity_test

import (
	"testing"

	"github.com/endorama/devid/internal/plugintest"
	"github.com/endorama/devid/plugins/identity"
	"github.com/stretchr/testify/assert"
)

func TestPlugin_Render(t *testing.T) {
	p := plugintest.GetPersona(t)

	i := identity.NewPlugin()
	i.LoadConfig(p.Config)

	r := i.Render(p.Name(), p.Location())

	expected := `export IDENTITY_EMAIL="bob@example.com"
export IDENTITY_NAME="Bob Testcase"
`
	assert.Equal(t, expected, r)
}
