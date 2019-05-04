package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(removeCmd)
}

var removeCmd = &cobra.Command{
	Use:   "rm",
	Short: "Delete all tasks on your to-do list",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		// Find all tasks
		tasks, err := taskManager.FindAll()
		if err != nil {
			exit(err)
		}
		// Delete all tasks from database
		for _, task := range tasks {
			err = taskManager.Delete(task)
			if err != nil {
				exit(err)
			}
		}
		err = os.Remove(dbPath)
		if err != nil {
			exit(err)
		}
		fmt.Println("Deleted all tasks on your to-do list")
	},
}
