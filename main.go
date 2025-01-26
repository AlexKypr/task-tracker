package main

import (
	"fmt"
	"os"

	cmd "github.com/AlexKypr/task-tracker/cmd/task-cli"
	"github.com/AlexKypr/task-tracker/internal/storage"
)

func main() {
	// Initialize the store (choose the storage implementation)
	store := storage.NewJSONStore("./tasks.json")

	// Create the CLI
	cli := cmd.NewCLI(store)

	// Execute the respective command on CLI
	if err := cli.Execute(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
