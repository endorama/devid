package aws_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

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

	expected := `export AWS_SHARED_CREDENTIALS_FILE=testdata/alice/aws/credentials
`
	assert.Equal(t, expected, r)
}

func TestPlugin_RenderWithLocalConfig(t *testing.T) {
	config := plugintest.GetConfig(t, "bob").Sub("aws")
	require.NotNil(t, config)

	p := plugintest.GetPersona(t, "bob")

	i := aws.NewPlugin()

	err := i.Configure(config)
	require.NoError(t, err)

	assert.True(t, plugintest.IsEnabled(t, "aws", p.Config), "plugin is not enabled for this persona")

	r := i.Render(p.Name(), p.Location())

	expected := `export AWS_CONFIG_FILE=testdata/bob/aws/config
export AWS_SHARED_CREDENTIALS_FILE=testdata/bob/aws/credentials
`
	assert.Equal(t, expected, r)
}

func TestPlugin_RenderWithCustomProfileName(t *testing.T) {
	config := plugintest.GetConfig(t, "charlie").Sub("aws")
	require.NotNil(t, config)

	p := plugintest.GetPersona(t, "charlie")

	i := aws.NewPlugin()

	err := i.Configure(config)
	require.NoError(t, err)

	assert.True(t, plugintest.IsEnabled(t, "aws", p.Config), "plugin is not enabled for this persona")

	r := i.Render(p.Name(), p.Location())

	expected := `export AWS_PROFILE="foobar"
export AWS_SHARED_CREDENTIALS_FILE=testdata/charlie/aws/credentials
`
	assert.Equal(t, expected, r)
}

func TestPlugin_RenderWithCustomProfileNamePersona(t *testing.T) {
	config := plugintest.GetConfig(t, "dan").Sub("aws")
	require.NotNil(t, config)

	p := plugintest.GetPersona(t, "dan")

	i := aws.NewPlugin()

	err := i.Configure(config)
	require.NoError(t, err)

	assert.True(t, plugintest.IsEnabled(t, "aws", p.Config), "plugin is not enabled for this persona")

	r := i.Render(p.Name(), p.Location())

	expected := `export AWS_PROFILE="dan"
export AWS_SHARED_CREDENTIALS_FILE=testdata/dan/aws/credentials
`
	assert.Equal(t, expected, r)
}
