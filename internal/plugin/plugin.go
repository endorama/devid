package plugin

import "github.com/spf13/viper"

// Pluggable interface allow a plugin to be leveraged.
// Is the common interface implemented by all plugins.
type Pluggable interface {
	Name() string
}

// Configurable interface allow a plugin to load configuration from the profile
// folder.
type Configurable interface {
	Configure(v *viper.Viper) error
}

// Generated allow a plugin implementing the Generator interface to return a
// list of files to be generated.
// File generation is then handled by the plugin/manager.
// Should contain maps of <file name>: <file content>
type Generated struct {
	Executables []GeneratedFile
	Files       []GeneratedFile
}

type GeneratedFile struct {
	Name    string
	Content string
}

// Generator interface allow a plugin to generate content before rendering.
// Generation has side effects, so is required to provide a method to cleanup
// generated content to allow for idempotency.
type Generator interface {
	Generate(location string) (Generated, error)
}

// Renderable interface allow a plugin to render content in the shell loader
// and runner files.
type Renderable interface {
	Render(name, location string) string
}

// Setuppable interface allow a plugin to perform setup steps before rendering.
type Setuppable interface {
	Setup(name string) error
}

// PluggableInstantiator is a proxy type for the init function for a plugin.
type PluggableInstantiator func() Pluggable
