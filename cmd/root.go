package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/jace-ys/taskar/pkg/manager"
	"github.com/jace-ys/taskar/pkg/task"
)

var taskFile = "data/tasks.db"
var taskManager *manager.TaskManager
var err error

var rootCmd = &cobra.Command{
	Use:     "taskar",
	Short:   "Taskar is a CLI for managing your to-do list",
	Version: "0.1",
}

func Execute() {
	// Create taskManager and setup database
	taskManager, err = manager.New(taskFile)
	if err != nil {
		exit(err)
	}
	err = rootCmd.Execute()
	if err != nil {
		exit(err)
	}
}

func exit(err error) {
	fmt.Println(err.Error())
	os.Exit(1)
}

// Return an error if provided task is nil
func validateTask(t *task.Task) {
	if t == nil {
		exit(fmt.Errorf("Requested task could not be found"))
	}
}
