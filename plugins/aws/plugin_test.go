package aws_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/endorama/devid/plugins/name"
)

func TestNewPlugin(t *testing.T) {
	p := name.NewPlugin()
	assert.IsType(t, &name.Plugin{}, p)
}
