package backup

import (
	"os"
)

// NewTask initialize a Task
func NewTask(name, source string, destination *os.File) (Task, error) {
	return Task{name, source, destination}, nil
}
