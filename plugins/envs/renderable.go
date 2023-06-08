package envs

import (
	"fmt"
	"strings"
)

func (p Plugin) Render(_, _ string) string {
	sb := strings.Builder{}

	for name, value := range p.config {
		name = strings.ToUpper(name)
		sb.WriteString(fmt.Sprintf("export %s=\"%s\"\n", name, value))
	}

	return sb.String()
}
