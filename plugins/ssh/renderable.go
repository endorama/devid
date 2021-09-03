package ssh

import (
	"fmt"
	"log"
	"path"
	"strings"
	"text/template"
)

const t = `# create agent cache if missing
if [ ! -f {{.CachePath}} ]; then
	ssh-agent -s | sed "s/echo/# echo/" > {{.CachePath}}
	chown "$USER:$USER" {{.CachePath}}
	chmod 600 {{.CachePath}}
fi
# load agent
source {{.CachePath}}
# add ssh keys, if not already loaded
{{ range $key, $value := .Keys -}}
if ! ssh-add -l 2> /dev/null | grep {{$value}} > /dev/null; then
	ssh-add {{$value}} > /dev/null
fi
{{end -}}`

// Render returns content rendered by the plugin.
// Implements `plugin.Renderable` interface.
func (p Plugin) Render(personaName, personaDirectory string) string {
	config := Config{}
	sb := strings.Builder{}
	sshPluginFolder := path.Join(personaDirectory, pluginName)

	config.CachePath = fmt.Sprintf(p.config.CachePath, personaName)

	// NOTE: to avoid specifying the entire path to the key, we expect them to
	// be in {{profileLocation}}/ssh
	for _, value := range p.config.Keys {
		// FIXME: prevent directory traversal
		config.Keys = append(config.Keys, path.Join(sshPluginFolder, value))
	}

	tpl, err := template.New("plugin-ssh").Parse(t)
	if err != nil {
		log.Fatal(err)
	}
	tpl.Execute(&sb, config)

	return sb.String()
}
