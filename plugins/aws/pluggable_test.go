package aws_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/endorama/devid/internal/plugin"
	"github.com/endorama/devid/plugins/name"
)

func TestPlugin_Pluggable(t *testing.T) {
	p := name.NewPlugin()
	assert.Implements(t, (*plugin.Pluggable)(nil), p)
}

func TestPlugin_Name(t *testing.T) {
	p := name.NewPlugin()
	assert.Equal(t, "name", p.Name())
}
