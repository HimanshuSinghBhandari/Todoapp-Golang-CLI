package cmd

import (
	"fmt"
	"strconv"
	"todo-cli/internal/todo"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use: "delete",
	Short: "Delete a todo by ID",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please peovide a task ID.")
			return 
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid task id");
			return
		}
		todo.DeleteTask(id)
        fmt.Printf("Deleted task with ID: %d",id);
	},
}