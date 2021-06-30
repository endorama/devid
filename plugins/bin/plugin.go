package bin

const pluginName = "bin"

type Plugin struct{}

func NewPlugin() *Plugin {
	return &Plugin{}
}

func (p Plugin) Name() string {
	return pluginName
}
