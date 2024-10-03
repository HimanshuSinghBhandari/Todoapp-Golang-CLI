package cmd

import (
	"todo-cli/internal/todo"

	"github.com/spf13/cobra"
)

var saveCmd = &cobra.Command{
	Use: "save",
	Short: "save todos to a file",
	Run: func(cmd *cobra.Command, args []string) {
		 todo.SaveAndLoadtasks()
	},
}