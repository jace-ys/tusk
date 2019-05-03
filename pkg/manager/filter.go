package manager

import (
	"strings"

	"github.com/jace-ys/taskar/pkg/task"
)

type FilterOptions struct {
	Name     string
	Category string
	DueDate  string
	Watching bool
}

func (tm *TaskManager) Filter(tasks task.TaskSlice, opts FilterOptions) task.TaskSlice {
	var filtered task.TaskSlice
	// Filter based on watch status
	if opts.Watching {
		for _, t := range tasks {
			if t.Watching {
				filtered = append(filtered, t)
			}
		}
		return filtered
	}
	// Filter using other parameters
	// Use a map of pointers to structs to hold unique values
	filterMap := make(map[*task.Task]struct{})
	for _, t := range tasks {
		if nameContains(t.Name, opts.Name) {
			filterMap[t] = struct{}{}
		}
		if strings.EqualFold(t.Category, opts.Category) {
			filterMap[t] = struct{}{}
		}
		if strings.EqualFold(t.DueDate, opts.DueDate) {
			filterMap[t] = struct{}{}
		}
	}
	// Convert back from map to task.TaskSlice
	for k, _ := range filterMap {
		filtered = append(filtered, k)
	}
	return filtered
}

// Check if str contains substr
func nameContains(str, substr string) bool {
	if substr != "" && strings.Contains(strings.ToLower(str), strings.ToLower(substr)) {
		return true
	}
	return false
}
