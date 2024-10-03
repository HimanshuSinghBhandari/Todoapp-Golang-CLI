package cmd

import (
	"todo-cli/internal/todo"

	"github.com/spf13/cobra"
)

// Tabular display command

var tableCmd = &cobra.Command{
	Use: "table",
	Short: "Display tools in a table format",
	Run: func(cmd *cobra.Command, args []string) {
		todo.DisplayTaskTable()
	},
}