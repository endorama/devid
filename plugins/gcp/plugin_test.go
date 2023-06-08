//go:build skip
// +build skip

package gcp_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/endorama/devid/plugins/gcp"
)

func TestNewPlugin(t *testing.T) {
	p := gcp.NewPlugin()
	assert.IsType(t, &gcp.Plugin{}, p)
}
