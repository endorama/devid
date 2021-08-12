package envs_test

import (
	"testing"

	"github.com/endorama/devid/internal/plugin"
	"github.com/endorama/devid/internal/plugintest"
	"github.com/endorama/devid/plugins/envs"
	"github.com/stretchr/testify/assert"
)

func TestPlugin_Renderable(t *testing.T) {
	p := envs.NewPlugin()
	assert.Implements(t, (*plugin.Renderable)(nil), p)
}

func TestPlugin_Render(t *testing.T) {
	p := plugintest.GetPersona(t)

	i := envs.NewPlugin()
	i.LoadConfig(p.Config)

	r := i.Render(p.Name(), p.Location())

	expected := `export FOO="bar"
`
	assert.Equal(t, expected, r)
}
