package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/jace-ys/taskar/pkg/printer"
)

func init() {
	listCmd.AddCommand(listCountCmd)
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display all tasks on your to-do list",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		// Find all tasks
		tasks, err := taskManager.FindAll()
		if err != nil {
			exit(err)
		}
		printer.PrintTable(tasks)
	},
}

var listCountCmd = &cobra.Command{
	Use:   "count",
	Short: "Count the number of tasks on your to-do list",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		// Find all tasks
		tasks, err := taskManager.FindAll()
		if err != nil {
			exit(err)
		}
		fmt.Println(len(tasks))
	},
}
