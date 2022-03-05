package envs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/endorama/devid/plugins/envs"
)

func TestNewPlugin(t *testing.T) {
	p := envs.NewPlugin()
	assert.IsType(t, &envs.Plugin{}, p)
}
