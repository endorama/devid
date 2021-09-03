package identity

import (
	"strings"
)

func (p Plugin) Render(profileName, profileLocation string) string {
	sb := strings.Builder{}

	sb.WriteString("export IDENTITY_EMAIL=\"" + p.config.Email + "\"\n")
	sb.WriteString("export IDENTITY_NAME=\"" + p.config.Name + "\"\n")

	return sb.String()
}
