package tmux

import "fmt"

// Render returns content rendered by the plugin.
// Implements `plugin.Renderable` interface.
func (p Plugin) Render(personaName, personaDirectory string) string {
	return fmt.Sprintf("export TMUX_SOCKET_NAME=\"%s.%s\"\n", pluginName, personaName)
}
