package cmd

import (
	"fmt"
	"todo-cli/internal/todo"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use: "add",
	Short: "Add a new Todo",
	Run: func (cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide a Todo Description.")
			return 
		}
		task:= args[0]
		todo.AddTask(task)
		fmt.Printf("Added Tasks: %s\n",task)
	},
}