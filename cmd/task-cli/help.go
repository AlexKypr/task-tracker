package cmd

import "fmt"

// showHelp displays the usage instructions and available commands for the task-cli.
func (c *CLI) showHelp() {
	fmt.Println(`
Usage:
    task-cli <command> [arguments]

Commands:
    add <task>                             Add a new task
                                            Example: task-cli add "Buy groceries"

    update <id> <updated task>             Update an existing task
                                            Example: task-cli update 1 "Buy groceries and cook dinner"

    delete <id>                            Delete a task by its ID
                                            Example: task-cli delete 1

    mark-in-progress <id>                  Mark a task as in progress
                                            Example: task-cli mark-in-progress 1

    mark-done <id>                         Mark a task as done
                                            Example: task-cli mark-done 1

    list                                   List all tasks
                                            Example: task-cli list

    list <status>                          List tasks by status (done, todo, in-progress)
                                            Example: task-cli list done
										
    help                                   Print this help message
                                            Example: task-cli help

Notes:
    - Task IDs are integers assigned sequentially when tasks are created.
    - Status values are case-insensitive: "done", "todo", "in-progress".
    - Use quotes for tasks with spaces.

For more help, visit the https://github.com/AlexKypr/task-tracker/blob/main/README.md.
`)
}
