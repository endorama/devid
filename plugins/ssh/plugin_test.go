package ssh_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/endorama/devid/plugins/ssh"
)

func TestNewPlugin(t *testing.T) {
	p := ssh.NewPlugin()
	assert.IsType(t, &ssh.Plugin{}, p)
}
