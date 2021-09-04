package plugin

/*
Generator interface allow a plugin to generate content before rendering.

Generation has side effects. To provide a way for the application to cleanup
generated artifacts, this method must returns a structure containing information
about the artifacts themselves. Filesystem operations are handled by the
application and must not be handled the plugin itself.

This allows to provide idempotent generation.
*/
type Generator interface {
	Generate(location string) (Generated, error)
}

/*
Generated allow a plugin implementing the Generator interface to return a
list of files to be generated.

File generation is then handled by the plugin/manager.

GeneratedFiles added to the Executables slice are made executable, readable and
writable by the current user.
GeneratedFiles added to the Files slices are written to disk readable and
writable only by the current user.
*/
type Generated struct {
	Executables []GeneratedFile
	Files       []GeneratedFile
}

// GeneratedFile represent a single file to be managed by plugin/manager.
type GeneratedFile struct {
	// Name of the file on disk
	Name string
	// Content of the file
	Content string
}
