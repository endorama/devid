package plugin

// Pluggable interface allow a plugin to be leveraged.
// Is the common interface implemented by all plugins.
type Pluggable interface {
	Name() string
}

// Configurable interface allow a plugin to load configuration from the profile
// folder.
type Configurable interface {
	Config() interface{}
	LoadConfig(configFile []byte) error
}

// Generator interface allow a plugin to generate content before rendering.
type Generator interface {
	Generate(location string) error
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
