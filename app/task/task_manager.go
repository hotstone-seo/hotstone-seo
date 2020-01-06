package task

import "go.uber.org/dig"

// Manager responsible manage the task
type Manager struct {
	dig.In
}

// Start the task
func (*Manager) Start() (err error) {
	// TODO:
	return
}
