package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"

	"github.com/jace-ys/tusk/pkg/manager"
	"github.com/jace-ys/tusk/pkg/task"
)

var taskFile = "go/data/tusk/tasks.db"
var dbPath string
var taskManager *manager.TaskManager

var rootCmd = &cobra.Command{
	Use:     "tusk",
	Short:   "Tusk is a CLI for managing your to-do list",
	Version: "0.1",
}

func Execute() {
	home, err := homedir.Dir()
	dbPath = filepath.Join(home, taskFile)
	// Create taskManager and setup database
	taskManager, err = manager.New(dbPath)
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
