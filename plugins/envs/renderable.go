package envs

import (
	"fmt"
	"strings"
)

func (p Plugin) Render(profileName, profileLocation string) string {
	config := p.Config().(Config)
	sb := strings.Builder{}
	for name, value := range config {
		name = strings.ToUpper(name)
		sb.WriteString(fmt.Sprintf("export %s=\"%s\"\n", name, value))
	}
	return sb.String()
}
