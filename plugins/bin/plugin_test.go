package bin_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/endorama/devid/plugins/bin"
)

func TestNewPlugin(t *testing.T) {
	p := bin.NewPlugin()
	assert.IsType(t, &bin.Plugin{}, p)
}
