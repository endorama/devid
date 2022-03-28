package gcp_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/endorama/devid/internal/plugin"
	"github.com/endorama/devid/plugins/gcp"
)

func TestPlugin_Pluggable(t *testing.T) {
	p := gcp.NewPlugin()
	assert.Implements(t, (*plugin.Pluggable)(nil), p)
}

func TestPlugin_Name(t *testing.T) {
	p := gcp.NewPlugin()
	assert.Equal(t, "gcp", p.Name())
}
