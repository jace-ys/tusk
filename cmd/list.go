package cmd

import (
	"github.com/spf13/cobra"

	"github.com/jace-ys/taskar/pkg/printer"
)

func init() {
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
