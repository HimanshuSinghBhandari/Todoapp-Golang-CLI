package cmd

import (
	"todo-cli/internal/todo"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use: "list",
	Short: "List of all todos",
	Run: func(cmd *cobra.Command, args []string) {
		todo.ListTasks()
	},
}