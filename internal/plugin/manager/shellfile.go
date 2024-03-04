package manager

import (
	// required by go:embed.
	_ "embed"
	"fmt"
	"log"
	"strings"
	"text/template"
	"time"

	"github.com/spf13/viper"

	"github.com/endorama/devid/internal/persona"
	"github.com/endorama/devid/internal/plugin"
)

//
//nolint:gochecknoglobals // required by embed
//go:embed load.sh.txt
var shellLoader string

// ShellLoader generate profile shell loader file.
func ShellLoader(p persona.Persona) (string, error) {
	// data is passed to the template
	data := struct {
		ActivePersona     string
		ActivePersonaPath string
		Name              string
		Date              string
		Location          string
		RenderedPlugins   string
		Shell             string
	}{
		ActivePersona:     viper.GetString("active_persona_env"),
		ActivePersonaPath: viper.GetString("active_persona_path_env"),
		Name:              p.Name(),
		Date:              time.Now().Format(time.RFC822),
		Location:          p.Location(),
		Shell:             viper.GetString("shell"),
	}

	log.SetPrefix("shell-loader-generator ")
	defer log.SetPrefix("")

	tmpl, err := template.New("shellLoaderFile").Parse(shellLoader)
	if err != nil {
		return "", fmt.Errorf("cannot create shellLoaderFile template: %w", err)
	}

	log.Printf("%+v", tmpl)

	logPlugins(plugins)

	data.RenderedPlugins = renderPlugins(p, plugins)

	content := strings.Builder{}
	if err = tmpl.Execute(&content, data); err != nil {
		return "", fmt.Errorf("cannot execute shellLoaderFile template: %w", err)
	}

	return content.String(), nil
}

func logPlugins(plugins []Plugin) {
	availablePlugins := strings.Builder{}
	enabledPlugins := strings.Builder{}

	for _, plg := range plugins {
		availablePlugins.Write([]byte(plg.Instance.Name()))

		if plg.Enabled {
			enabledPlugins.Write([]byte(plg.Instance.Name()))
		}

		if plg != plugins[len(plugins)-1] {
			availablePlugins.WriteByte(',')
			enabledPlugins.WriteByte(',')
		}
	}

	log.Printf("available plugins: %s\n", availablePlugins.String())
	log.Printf("enabled plugins  : %s\n", enabledPlugins.String())
}

func renderPlugins(p persona.Persona, plugins []Plugin) string {
	sb := strings.Builder{}

	for _, plg := range plugins {
		if renderablePlugin, ok := plg.Instance.(plugin.Renderable); ok {
			if plg.Enabled {
				log.Printf("rendering plugin: %s", plg.Instance.Name())
				sb.WriteString(fmt.Sprintf("# plugin %s\n", plg.Instance.Name()))
				sb.WriteString(renderablePlugin.Render(p.Name(), p.Location()))
			}
		}
	}

	return sb.String()
}
