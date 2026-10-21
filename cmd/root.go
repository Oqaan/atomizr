package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "atomizr",
	Short: "Suggests which fragrance to wear today",
}

// Execute runs the CLI and exits with status 1 on error.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
