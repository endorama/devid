package backup

import "io"

// NewTask initialize a Task.
// TODO: maybe rename to Prepare() or New()?
func NewTask(name, source string, out io.Writer) (Task, error) {
	return Task{name, source, out}, nil
}
