package bin

import "strings"

func (p Plugin) Render(_, profileLocation string) string {
	sb := strings.Builder{}
	sb.WriteString("export PATH=\"" + profileLocation + "/bin:$PATH\"\n")

	return sb.String()
}
