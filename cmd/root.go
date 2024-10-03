// # Main entry point for CLI
package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "todo",
	Short: "Todo CLI Application",
	Long: "A CLI Based TODO application for managing tasks",
}

func Execute() {
	if err := rootCmd.Execute(); err!= nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	 rootCmd.AddCommand(addCmd);
	 rootCmd.AddCommand(deleteCmd);
	 rootCmd.AddCommand(listCmd);
	 rootCmd.AddCommand(tableCmd);
	 rootCmd.AddCommand(saveCmd);
}