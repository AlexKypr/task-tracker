package models

// TaskCommands defines the available commands for the task CLI.
type TaskCommands string

const (
	// Commands for task management
	Add            TaskCommands = "add"
	Update         TaskCommands = "update"
	Delete         TaskCommands = "delete"
	MarkInProgress TaskCommands = "mark-in-progress"
	MarkDone       TaskCommands = "mark-done"
	List           TaskCommands = "list"
	Help           TaskCommands = "help"
)

func NewTaskCommand(input string) (TaskCommands, bool) {
	switch TaskCommands(input) {
	case Add, Update, Delete, MarkInProgress, MarkDone, List, Help:
		return TaskCommands(input), true
	default:
		return "", false
	}
}
