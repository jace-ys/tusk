package cmd

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/jace-ys/tusk/pkg/manager"
	"github.com/jace-ys/tusk/pkg/printer"
)

func init() {
	watchCmd.AddCommand(watchListCmd, watchCountCmd)
	rootCmd.AddCommand(watchCmd, unwatchCmd)
}

var watchCmd = &cobra.Command{
	Use:   "watch [task id]",
	Short: "Watch a task on your to-do list",
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
				// Set task's watch status to true
				err = taskManager.Update(task.Watch(true))
				if err != nil {
					exit(err)
				}
				fmt.Printf("Watching task \"%s\"\n", task.Name)
			}
		}
	},
}

var unwatchCmd = &cobra.Command{
	Use:   "unwatch [task id]",
	Short: "Unwatch a task on your to-do list",
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
				// Set task's watch status to false
				err = taskManager.Update(task.Watch(false))
				if err != nil {
					exit(err)
				}
				fmt.Printf("No longer watching task \"%s\"\n", task.Name)
			}
		}
	},
}

var watchListCmd = &cobra.Command{
	Use:   "list",
	Short: "Display all watched tasks on your to-do list",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		// Find all tasks
		tasks, err := taskManager.FindAll()
		if err != nil {
			exit(err)
		}
		// Filter tasks based on watch status
		filtered := taskManager.Filter(tasks, manager.FilterOptions{Watching: true})
		printer.PrintTable(filtered)
	},
}

var watchCountCmd = &cobra.Command{
	Use:   "count",
	Short: "Count the number of watched tasks on your to-do list",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		// Find all tasks
		tasks, err := taskManager.FindAll()
		if err != nil {
			exit(err)
		}
		// Filter tasks based on watch status
		filtered := taskManager.Filter(tasks, manager.FilterOptions{Watching: true})
		fmt.Println(len(filtered))
	},
}
