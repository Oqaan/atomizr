package cmd

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/Oqaan/atomizr/internal/collection"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all fragrances in your collection",
	RunE: func(*cobra.Command, []string) error {
		configDir, err := os.UserConfigDir()
		if err != nil {
			return fmt.Errorf("finding config directory: %w", err)
		}
		path := filepath.Join(configDir, "atomizr", "collection.yaml")
		fragrances, err := collection.Load(path)
		if err != nil {
			if errors.Is(err, fs.ErrNotExist) {
				return fmt.Errorf("no collection found at %s\nCreate this file to add your fragrances.", path)
			}
			return err
		}
		fmt.Printf("Found %d fragrances.\n", len(fragrances))
		return nil
	},
}

// init registers the list command with the root command.
func init() {
	rootCmd.AddCommand(listCmd)
}
