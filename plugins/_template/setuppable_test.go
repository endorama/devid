package name_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/endorama/devid/internal/plugin"
	"github.com/endorama/devid/plugins/name"
)

func TestPlugin_Setuppable(t *testing.T) {
	p := name.NewPlugin()
	assert.Implements(t, (*plugin.Setuppable)(nil), p)
}

func TestPlugin_Setup(t *testing.T) {
	// appFS := afero.NewMemMapFs()
	// p := plugintest.GetPersona(t)
	//
	// plg := name.TestNewPlugin(appFS)
	//
	// err := plg.Setup(p.Location())
	// assert.Nil(t, err)

	t.Skip("Not implemented")
}
