package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all fragrances in your collection",
	Run: func(*cobra.Command, []string) {
		fmt.Println("No fragrances yet.")
	},
}

// init registers the list command with the root command.
func init() {
	rootCmd.AddCommand(listCmd)
}
