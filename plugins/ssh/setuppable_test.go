package ssh_test

import (
	"path"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"

	"github.com/endorama/devid/internal/plugin"
	"github.com/endorama/devid/internal/plugintest"
	"github.com/endorama/devid/plugins/ssh"
)

func TestPlugin_Setuppable(t *testing.T) {
	p := ssh.NewPlugin()
	assert.Implements(t, (*plugin.Setuppable)(nil), p)
}

func TestPlugin_Setup(t *testing.T) {
	appFS := afero.NewMemMapFs()
	p := plugintest.GetPersona(t, "alice")

	plg := ssh.TestNewPlugin(appFS)

	err := plg.Setup(p.Location())
	assert.Nil(t, err)

	s, err := appFS.Stat(path.Join(p.Location(), "ssh"))
	assert.Nil(t, err)
	assert.True(t, s.IsDir())

	assert.EqualValues(t, s.Mode().Perm(), int(0750))
}
