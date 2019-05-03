package task

import (
	"strings"
	"time"
)

var timeFormat = "2/1/2006"

type TaskSlice []*Task

type Task struct {
	ID       string
	Name     string
	Category string
	DueDate  string
	Comment  string
	Watching bool
}

func New(category string, task string) *Task {
	return &Task{
		Name:     task,
		Category: category,
	}
}

func (t *Task) SetCategory(category string) *Task {
	t.Category = category
	return t
}

func (t *Task) SetName(name string) *Task {
	t.Name = name
	return t
}

func (t *Task) SetDue(date string) *Task {
	// Parse the provided date
	d, err := time.Parse(timeFormat, date)
	switch {
	// Set DueDate to tomorrow's date if `tomorrow` specified
	case strings.EqualFold(date, "tomorrow"):
		t.DueDate = time.Now().AddDate(0, 0, 1).Format(timeFormat)
	// Set DueDate to "-" if error occured in parsing the provided date
	case err != nil:
		t.DueDate = "-"
	// Otherwise, set DueDate to provided date
	default:
		t.DueDate = d.Format(timeFormat)
	}
	return t
}

func (t *Task) SetComment(comment string) *Task {
	t.Comment = comment
	return t
}

func (t *Task) Watch(watch bool) *Task {
	t.Watching = watch
	return t
}
