package name

import "strings"

// Render returns content rendered by the plugin.
// Implements `plugin.Renderable` interface.
func (p Plugin) Render(personaName, personaDirectory string) string {
	// TODO: implement rendering logic
	sb := strings.Builder{}
	// sb.WriteString("plugin rendered content")

	return sb.String()
}
