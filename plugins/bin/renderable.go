package bin

import "strings"

func (p Plugin) Render(profileName, profileLocation string) string {
	sb := strings.Builder{}
	sb.WriteString("export PATH=\"" + profileLocation + "/bin:$PATH\"\n")

	return sb.String()
}
