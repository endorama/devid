package plugin

/*
Renderable interface allow a plugin to render content in the shell loader
and runner files.
*/
type Renderable interface {
	// Render should use internal state of the plugin to output a string to be
	// included in the Persona shell loader file.
	Render(name, location string) string
}
