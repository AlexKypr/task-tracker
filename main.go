package main

import (
	"fmt"
	"os"

	cmd "github.com/AlexKypr/task-tracker/cmd/task-cli"
)

func main() {
	if err := cmd.Execute(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
