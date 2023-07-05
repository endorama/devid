package aws_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/endorama/devid/internal/plugin"
	"github.com/endorama/devid/internal/plugintest"
	"github.com/endorama/devid/plugins/aws"
)

func TestPlugin_Renderable(t *testing.T) {
	p := aws.NewPlugin()
	assert.Implements(t, (*plugin.Renderable)(nil), p)
}

func TestPlugin_Render(t *testing.T) {
	p := plugintest.GetPersona(t, "alice")

	i := aws.NewPlugin()

	assert.True(t, plugintest.IsEnabled(t, "aws", p.Config), "plugin is not enabled for this persona")

	r := i.Render(p.Name(), p.Location())

	expected := `export AWS_CONFIG_FILE=testdata/alice/aws/config
export AWS_PROFILE="alice"
export AWS_SHARED_CREDENTIALS_FILE=testdata/alice/aws/credentials
`
	assert.Equal(t, expected, r)
}
