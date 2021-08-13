package ssh

import "github.com/endorama/devid/internal/plugin"

// Generate a file in the bin plugin directory.
// Implements the Generator interface.
func (p *Plugin) Generate(personaDirectory string) (plugin.Generated, error) {
	// TODO: implement file generation
	// file := strings.Builder{}

	return plugin.Generated{
		Executables: []plugin.GeneratedFile{},
		Files:       []plugin.GeneratedFile{},
	}, nil
}
