package aws

import "strings"

// Render returns content rendered by the plugin.
// Implements `plugin.Renderable` interface.
func (p *Plugin) Render(personaName, personaDirectory string) string {
	sb := strings.Builder{}

	if p.config.LocalConfig {
		sb.WriteString("export AWS_CONFIG_FILE=" + personaDirectory + "/aws/config\n")
	}

	sb.WriteString("export AWS_PROFILE=\"" + personaName + "\"\n")
	sb.WriteString("export AWS_SHARED_CREDENTIALS_FILE=" + personaDirectory + "/aws/credentials\n")

	return sb.String()
}
