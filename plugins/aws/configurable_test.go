package aws_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/endorama/devid/internal/plugin"
	"github.com/endorama/devid/internal/plugintest"
	"github.com/endorama/devid/plugins/aws"
)

func TestPlugin_Configurable(t *testing.T) {
	p := aws.NewPlugin()
	assert.Implements(t, (*plugin.Configurable)(nil), p)
}

func TestPlugin_Configure(t *testing.T) {
	config := plugintest.GetConfig(t, "alice").Sub("aws")
	assert.NotNil(t, config)

	p := aws.NewPlugin()

	err := p.Configure(config)
	if err != nil {
		t.Errorf(fmt.Errorf("cannot load plugin config: %w", err).Error())
	}

	assert.Nil(t, err, "err should be nil")
}

func TestPlugin_ConfigureLocalConfig(t *testing.T) {
	config := plugintest.GetConfig(t, "bob").Sub("aws")
	require.NotNil(t, config)

	p := aws.NewPlugin()

	err := p.Configure(config)
	if err != nil {
		t.Errorf(fmt.Errorf("cannot load plugin config: %w", err).Error())
	}

	require.Nil(t, err, "err should be nil")
	require.True(t, p.Config().LocalConfig)
}
