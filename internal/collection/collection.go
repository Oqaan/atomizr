package collection

import (
	"fmt"
	"os"

	"go.yaml.in/yaml/v3"
)

// Fragrance is a single fragrance in the user's collection.
type Fragrance struct {
	Name          string     `yaml:"name"`
	Brand         string     `yaml:"brand"`
	Family        string     `yaml:"family"`
	Concentration string     `yaml:"concentration"`
	Intensity     Intensity  `yaml:"intensity"`
	Seasons       []Season   `yaml:"seasons"`
	Occasions     []Occasion `yaml:"occasions"`
}

// Load reads and parses the collection file at path.
func Load(path string) ([]Fragrance, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading collection: %w", err)
	}
	var fragrances []Fragrance
	err = yaml.Unmarshal(data, &fragrances)
	if err != nil {
		return nil, fmt.Errorf("parsing collection: %w", err)
	}
	// Check required fields. UnmarshalYAML is not called for keys that
	// are missing entirely, so they show up as zero values here.
	for _, f := range fragrances {
		if f.Intensity == 0 {
			return nil, fmt.Errorf("fragrance %q: intensity is missing", f.Name)
		}
		if len(f.Seasons) == 0 {
			return nil, fmt.Errorf("fragrance %q: seasons are missing", f.Name)
		}
		if len(f.Occasions) == 0 {
			return nil, fmt.Errorf("fragrance %q: occasions are missing", f.Name)
		}
	}
	return fragrances, nil
}
