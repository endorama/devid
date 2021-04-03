package identity

const PluginName = "identity"

func NewPlugin() *Plugin {
	return &Plugin{}
}

type Plugin struct {
	config Config
}

func (p Plugin) Name() string {
	return PluginName
}
