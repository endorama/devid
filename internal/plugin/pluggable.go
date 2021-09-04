package plugin

/*
Pluggable interface allow a plugin to be leveraged.

This interface must be implemented by all plugins.
*/
type Pluggable interface {
	// Name must returns the name of the plugin
	Name() string
}
