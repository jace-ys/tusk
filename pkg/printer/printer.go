package printer

import (
	"os"

	"github.com/olekukonko/tablewriter"

	"github.com/jace-ys/tusk/pkg/task"
)

func PrintTable(tasks task.TaskSlice) {
	// Format table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Task", "Category", "Due Date", "Comment", "Watching"})
	table.SetBorders(tablewriter.Border{Left: false, Top: true, Right: false, Bottom: true})
	for _, t := range tasks {
		// Create string for watch status
		var watchStatus string
		if t.Watching {
			watchStatus = "\u2713"
		} else {
			watchStatus = "-"
		}
		table.Append([]string{t.ID, t.Name, t.Category, t.DueDate, t.Comment, watchStatus})
	}
	table.Render()
}
