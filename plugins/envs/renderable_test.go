package envs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/endorama/devid/internal/plugin"
	"github.com/endorama/devid/internal/plugintest"
	"github.com/endorama/devid/plugins/envs"
)

func TestPlugin_Renderable(t *testing.T) {
	p := envs.NewPlugin()
	assert.Implements(t, (*plugin.Renderable)(nil), p)
}

func TestPlugin_Render(t *testing.T) {
	p := plugintest.GetPersona(t, "alice")
	cfg := p.Config.Sub("envs")
	assert.NotNil(t, cfg)

	i := envs.NewPlugin()
	if err := i.Configure(cfg); err != nil {
		panic(err)
	}

	r := i.Render(p.Name(), p.Location())

	expected := `export FOO="bar"
`
	assert.Equal(t, expected, r)
}
