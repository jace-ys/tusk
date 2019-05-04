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
		// Add task to filterMap if it passes all filters
		keep := []bool{
			ifContains(t.Name, opts.Name),
			ifExact(t.Category, opts.Category),
			ifExact(t.DueDate, opts.DueDate),
		}
		if allTrue(keep) {
			filterMap[t] = struct{}{}
		}
	}
	// Convert back from map to task.TaskSlice
	for k, _ := range filterMap {
		filtered = append(filtered, k)
	}
	return filtered
}

func allTrue(bools []bool) bool {
	for _, v := range bools {
		if !v {
			return false
		}
	}
	return true
}

func ifContains(str, substr string) bool {
	switch {
	// Return true if substr equals empty string
	case substr == "":
		fallthrough
	case strings.Contains(strings.ToLower(str), strings.ToLower(substr)):
		return true
	default:
		return false
	}
}

func ifExact(str, substr string) bool {
	switch {
	// Return true if substr equals empty string
	case substr == "":
		fallthrough
	case strings.EqualFold(str, substr):
		return true
	default:
		return false
	}
}
