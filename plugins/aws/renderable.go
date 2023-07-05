package aws

import "strings"

// Render returns content rendered by the plugin.
// Implements `plugin.Renderable` interface.
func (p Plugin) Render(_, personaDirectory string) string {
	sb := strings.Builder{}
	sb.WriteString("export AWS_CONFIG_FILE=" + personaDirectory + "/aws/config\n")
	sb.WriteString("export AWS_SHARED_CREDENTIALS_FILE=" + personaDirectory + "/aws/credentials\n")

	return sb.String()
}
