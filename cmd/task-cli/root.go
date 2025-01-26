package cmd

import (
	"fmt"

	"github.com/AlexKypr/task-tracker/internal/storage"
	"github.com/AlexKypr/task-tracker/models"
)

type CLI struct {
	store storage.Store
}

func NewCLI(store storage.Store) *CLI {
	return &CLI{
		store: store,
	}
}

// Execute is responsible to execute the supported commands of CLI
func (c *CLI) Execute(args []string) error {
	if len(args) < 2 {
		c.showHelp()
		return nil
	}

	command, ok := models.NewTaskCommand(args[1])
	if !ok {
		c.showHelp()
		return fmt.Errorf("command %v is not supported", args[1])
	}

	commandHandlers := map[models.TaskCommands]func([]string) error{
		models.Add:            c.addTask,
		models.Update:         c.updateTask,
		models.Delete:         c.deleteTask,
		models.MarkInProgress: c.markInProgress,
		models.MarkDone:       c.markDone,
		models.List:           c.listTasks,
		models.Help: func(_ []string) error {
			c.showHelp()
			return nil
		},
	}

	handler, ok := commandHandlers[command]
	if !ok {
		c.showHelp()
		return fmt.Errorf("command %s is valid but not yet implemented", command)
	}

	return handler(args[2:])
}
