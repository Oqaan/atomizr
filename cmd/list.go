package cmd

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/tabwriter"

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
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintf(w, "NAME\tBRAND\tFAMILY\tINTENSITY\tSEASONS\tOCCASIONS\n")
		for _, f := range fragrances {
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n", f.Name, f.Brand, f.Family, f.Intensity, strings.Join(f.Seasons, ", "), joinOccasions(f.Occasions))
		}
		return w.Flush()
	},
}

// init registers the list command with the root command.
func init() {
	rootCmd.AddCommand(listCmd)
}

// joinOccasions formats occasions as "office, casual".
// strings.Join only accepts []string, so each occasion is converted first.
func joinOccasions(occasions []collection.Occasion) string {
	var names []string
	for _, o := range occasions {
		names = append(names, o.String())
	}
	return strings.Join(names, ", ")
}
