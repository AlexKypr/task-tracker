package cmd

import (
	"fmt"

	"github.com/AlexKypr/task-tracker/models"
)

// Execute is responsible to execute the supported commands of CLI
func Execute(args []string) error {
	if len(args) < 2 {
		showHelp()
		return nil
	}

	command, ok := models.NewTaskCommand(args[1])
	if !ok {
		showHelp()
		return fmt.Errorf("command %v is not supported", args[1])
	}

	commandHandlers := map[models.TaskCommands]func([]string) error{
		models.Add:            addTask,
		models.Update:         updateTask,
		models.Delete:         deleteTask,
		models.MarkInProgress: markInProgress,
		models.MarkDone:       markDone,
		models.List:           listTasks,
		models.Help: func(_ []string) error {
			showHelp()
			return nil
		},
	}

	handler, ok := commandHandlers[command]
	if !ok {
		showHelp()
		return fmt.Errorf("command %s is valid but not yet implemented", command)
	}

	return handler(args[2:])
}
