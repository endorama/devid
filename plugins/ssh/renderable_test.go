package ssh_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/endorama/devid/internal/plugin"
	"github.com/endorama/devid/internal/plugintest"
	"github.com/endorama/devid/plugins/ssh"
)

func TestPlugin_Renderable(t *testing.T) {
	p := ssh.NewPlugin()
	assert.Implements(t, (*plugin.Renderable)(nil), p)
}

func TestPlugin_Render(t *testing.T) {
	p := plugintest.GetPersona(t, "alice")
	cfg := p.Config.Sub("ssh")
	assert.NotNil(t, cfg)

	i := ssh.NewPlugin()
	i.Configure(cfg)

	r := i.Render(p.Name(), p.Location())

	expected := `# create agent cache if missing
if [ ! -f /tmp/devid-alice-ssh-agent.tmp ]; then
	ssh-agent -s | sed "s/echo/# echo/" > /tmp/devid-alice-ssh-agent.tmp
	chown "$USER:$USER" /tmp/devid-alice-ssh-agent.tmp
	chmod 600 /tmp/devid-alice-ssh-agent.tmp
fi
# load agent
source /tmp/devid-alice-ssh-agent.tmp
# add ssh keys, if not already loaded
if ! ssh-add -l 2> /dev/null | grep testdata/alice/ssh/id_rsa > /dev/null; then
	ssh-add testdata/alice/ssh/id_rsa > /dev/null
fi
`
	assert.Equal(t, expected, r)
}
