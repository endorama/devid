package aws

const pluginName = "aws"

type Plugin struct{}

// NewPlugin instantiate a Plugin instance.
func NewPlugin() *Plugin {
	return &Plugin{}
}
