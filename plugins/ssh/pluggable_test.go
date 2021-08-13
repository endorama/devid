package ssh_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/endorama/devid/internal/plugin"
	"github.com/endorama/devid/plugins/ssh"
)

func TestPlugin_Pluggable(t *testing.T) {
	p := ssh.NewPlugin()
	assert.Implements(t, (*plugin.Pluggable)(nil), p)
}

func TestPlugin_Name(t *testing.T) {
	p := ssh.NewPlugin()
	assert.Equal(t, "ssh", p.Name())
}
