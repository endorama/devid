package bin_test

import (
	"testing"

	"github.com/endorama/devid/plugins/bin"
	"github.com/stretchr/testify/assert"
)

func TestNewPlugin(t *testing.T) {
	p := bin.NewPlugin()
	assert.IsType(t, &bin.Plugin{}, p)
}
