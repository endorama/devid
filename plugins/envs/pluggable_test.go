package envs_test

import (
	"testing"

	"github.com/endorama/devid/internal/plugin"
	"github.com/endorama/devid/plugins/envs"
	"github.com/stretchr/testify/assert"
)

func TestPlugin_Pluggable(t *testing.T) {
	p := envs.NewPlugin()
	assert.Implements(t, (*plugin.Pluggable)(nil), p)
}

func TestPlugin_Name(t *testing.T) {
	p := envs.NewPlugin()
	assert.Equal(t, "envs", p.Name())
}
