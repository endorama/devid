package bin_test

import (
	"testing"

	"github.com/endorama/devid/internal/plugin"
	"github.com/endorama/devid/plugins/bin"
	"github.com/stretchr/testify/assert"
)

func TestPlugin_Pluggable(t *testing.T) {
	p := bin.NewPlugin()
	assert.Implements(t, (*plugin.Pluggable)(nil), p)
}

func TestPlugin_Name(t *testing.T) {
	p := bin.NewPlugin()
	assert.Equal(t, "bin", p.Name())
}
