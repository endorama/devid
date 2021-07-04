package manager

import (
	// required by go:embed.
	_ "embed"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/endorama/devid/internal/persona"
	"github.com/endorama/devid/internal/plugin"
	"github.com/endorama/devid/internal/settings"
)

//go:embed load.sh.txt
// nolint:gochecknoglobals // required by embed
var shellLoader string

// ShellLoader generate profile shell loader file.
func ShellLoader(p persona.Persona) (string, error) {
	data := struct {
		ActivePersona     string
		ActivePersonaPath string
		Name              string
		Date              string
		Location          string
		RenderedPlugins   string
		Shell             string
	}{
		ActivePersona:     settings.ActivePersonaEnv,
		ActivePersonaPath: settings.ActivePersonaPathEnv,
		Name:              p.Name(),
		Date:              time.Now().Format(time.RFC822),
		Location:          p.Location(),
		Shell:             os.Getenv("SHELL"),
	}

	log.SetPrefix("shell-loader-generator ")
	defer log.SetPrefix("")

	tmpl, err := template.New("shellLoaderFile").Parse(shellLoader)
	if err != nil {
		return "", fmt.Errorf("cannot create shellLoaderFile template: %w", err)
	}

	log.Printf("%+v", tmpl)
	log.Printf("%+v", plugins)

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

	data.RenderedPlugins = sb.String()
	content := strings.Builder{}

	err = tmpl.Execute(&content, data)
	if err != nil {
		return "", fmt.Errorf("cannot execute shellLoaderFile template: %w", err)
	}

	return content.String(), nil
}
