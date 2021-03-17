package plugin

// Pluggable interface allow a plugin to be leveraged.
// Is the common interface implemented by all plugins.
type Pluggable interface {
	Name() string
	Renderable
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

// Configurable interface allow a plugin to load configuration from the profile
// folder.
type Configurable interface {
	Config() interface{}
	ConfigFile(location string) string
	LoadConfig(location string) error
}

// Generator interface allow a plugin to generate content before rendering.
type Generator interface {
	Generate(location string) error
}
