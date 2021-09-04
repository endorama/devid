package plugin

/*
Setuppable interface allow a plugin to perform setup steps before rendering.

Setup has side effects. Its purpose is to interact with the host system to
prepare it for supporting the Persona shell loading.
*/
type Setuppable interface {
	Setup(name string) error
}
