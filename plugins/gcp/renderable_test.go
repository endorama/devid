//go:build skip
// +build skip

package gcp_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/endorama/devid/internal/plugin"
	"github.com/endorama/devid/internal/plugintest"
	"github.com/endorama/devid/plugins/gcp"
)

func TestPlugin_Renderable(t *testing.T) {
	p := gcp.NewPlugin()
	assert.Implements(t, (*plugin.Renderable)(nil), p)
}

func TestPlugin_Render(t *testing.T) {
	p := plugintest.GetPersona(t)

	i := gcp.NewPlugin()
	i.LoadConfig(p.Config)

	r := i.Render(p.Name(), p.Location())

	expected := `not implemented
`
	assert.Equal(t, expected, r)
}
