package envs

const pluginName = "envs"

type Plugin struct {
	config Config
}

func NewPlugin() *Plugin {
	return &Plugin{}
}
