package models

// TaskStatus defines the available statuses of a task
type TaskStatus string

const (
	// Task statuses'
	Todo       TaskStatus = "todo"
	InProgress TaskStatus = "in-progress"
	Done       TaskStatus = "done"
)
