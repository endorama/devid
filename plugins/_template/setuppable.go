package name

// Setup allow the plugin to perform any require setup actions.
// Implements `plugin.Setuppable` interface.
func (p Plugin) Setup(personaDirectory string) error {
	return nil
}
