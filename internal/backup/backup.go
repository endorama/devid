package backup

import "io"

// NewTask initialize a Task.
func NewTask(name, source string, out io.Writer) (Task, error) {
	return Task{name, source, out}, nil
}
