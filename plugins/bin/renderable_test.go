package bin_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/endorama/devid/internal/plugin"
	"github.com/endorama/devid/internal/plugintest"
	"github.com/endorama/devid/plugins/bin"
)

func TestPlugin_Renderable(t *testing.T) {
	p := bin.NewPlugin()
	assert.Implements(t, (*plugin.Renderable)(nil), p)
}

func TestPlugin_Render(t *testing.T) {
	p := plugintest.GetPersona(t, "alice")

	i := bin.NewPlugin()

	r := i.Render(p.Name(), p.Location())

	expected := fmt.Sprintf(`export PATH="%s/bin:$PATH"
`, p.Location())
	assert.Equal(t, expected, r)
}
