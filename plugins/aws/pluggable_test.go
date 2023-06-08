package aws_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/endorama/devid/internal/plugin"
	"github.com/endorama/devid/plugins/aws"
)

func TestPlugin_Pluggable(t *testing.T) {
	p := aws.NewPlugin()
	assert.Implements(t, (*plugin.Pluggable)(nil), p)
}

func TestPlugin_Name(t *testing.T) {
	p := aws.NewPlugin()
	assert.Equal(t, "aws", p.Name())
}
