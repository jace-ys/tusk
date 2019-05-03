package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(doneCmd)
}

var doneCmd = &cobra.Command{
	Use:   "done [task id]",
	Short: "Mark a task on your to-do list as complete",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Loop through provided ID's
		for _, arg := range args {
			taskId := arg
			// Find task with ID
			task, err := taskManager.FindOne(taskId)
			if err != nil {
				exit(err)
			}
			// Display error message if no task found
			if task == nil {
				fmt.Printf("No task with ID \"%s\" could be found\n", taskId)
			} else {
				// Delete task from database
				err = taskManager.Delete(task)
				if err != nil {
					exit(err)
				}
				fmt.Printf("Marked task \"%s\" as complete\n", task.Name)
			}
		}
	},
}
