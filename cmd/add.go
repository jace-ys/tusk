package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/jace-ys/taskar/pkg/task"
)

var addFlags struct {
	category string
	dueDate  string
	comment  string
	watch    bool
}

func init() {
	addCmd.Flags().StringVarP(&addFlags.category, "tag", "t", "general", "Tag a category to the given task")
	addCmd.Flags().StringVarP(&addFlags.dueDate, "due", "d", "-", "Set a due date for the given task")
	addCmd.Flags().StringVarP(&addFlags.comment, "msg", "m", "-", "Add a comment message for the given task")
	addCmd.Flags().BoolVarP(&addFlags.watch, "watch", "w", false, "Watch the given task")
	rootCmd.AddCommand(addCmd, amendCmd, tagCmd, setDueCmd, commentCmd)
}

var addCmd = &cobra.Command{
	Use:   "add [task name]",
	Short: "Add a new task to your to-do list",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Create new task using provided flags and arguments
		taskName := strings.Join(args, " ")
		t := task.New(addFlags.category, taskName).SetComment(addFlags.comment).SetDue(addFlags.dueDate).Watch(addFlags.watch)
		// Add new task to database
		err := taskManager.Create(t)
		if err != nil {
			exit(err)
		}
		fmt.Printf("Added \"%s\" to your to-do list\n", t.Name)
	},
}

var amendCmd = &cobra.Command{
	Use:   "amend [task id] [new task]",
	Short: "Amend the name of a task on your to-do list",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		taskId := args[0]
		t, err := taskManager.FindOne(taskId)
		if err != nil {
			exit(err)
		}
		validateTask(t)
		name := strings.Join(args[1:], " ")
		err = taskManager.Update(t.SetName(name))
		if err != nil {
			exit(err)
		}
		fmt.Printf("Amended task to \"%s\"\n", t.Name)
	},
}

var tagCmd = &cobra.Command{
	Use:   "tag [task id] [category]",
	Short: "Tag a category to a task on your to-do list",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		taskId := args[0]
		t, err := taskManager.FindOne(taskId)
		if err != nil {
			exit(err)
		}
		validateTask(t)
		category := strings.Join(args[1:], " ")
		err = taskManager.Update(t.SetCategory(category))
		if err != nil {
			exit(err)
		}
		fmt.Printf("Tagged category \"%s\" to task \"%s\"\n", t.Category, t.Name)
	},
}

var setDueCmd = &cobra.Command{
	Use:   "set-due [task id] [date]",
	Short: "Set a due date for a task on your to-do list",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		taskId := args[0]
		t, err := taskManager.FindOne(taskId)
		if err != nil {
			exit(err)
		}
		validateTask(t)
		date := strings.Join(args[1:], " ")
		err = taskManager.Update(t.SetDue(date))
		if err != nil {
			exit(err)
		}
		fmt.Printf("Set due date \"%s\" to task \"%s\"\n", t.DueDate, t.Name)
	},
}

var commentCmd = &cobra.Command{
	Use:   "comment [task id] [message]",
	Short: "Add a comment message to a task on your to-do list",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		taskId := args[0]
		t, err := taskManager.FindOne(taskId)
		if err != nil {
			exit(err)
		}
		validateTask(t)
		comment := strings.Join(args[1:], " ")
		err = taskManager.Update(t.SetComment(comment))
		if err != nil {
			exit(err)
		}
		fmt.Printf("Added comment \"%s\" to task \"%s\"\n", t.Comment, t.Name)
	},
}
