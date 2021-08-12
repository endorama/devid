package envs_test

import (
	"testing"

	"github.com/endorama/devid/plugins/envs"
	"github.com/stretchr/testify/assert"
)

func TestNewPlugin(t *testing.T) {
	p := envs.NewPlugin()
	assert.IsType(t, &envs.Plugin{}, p)
}
