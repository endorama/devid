package aws_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/endorama/devid/plugins/aws"
)

func TestNewPlugin(t *testing.T) {
	p := aws.NewPlugin()
	assert.IsType(t, &aws.Plugin{}, p)
}
