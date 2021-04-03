package identity

import "strings"

func (p Plugin) Render(profileName, profileLocation string) string {
	config := p.Config().(Config)
	sb := strings.Builder{}

	sb.WriteString("export IDENTITY_EMAIL=\"" + config.Email + "\"\n")
	sb.WriteString("export IDENTITY_NAME=\"" + config.Name + "\"\n")

	return sb.String()
}
