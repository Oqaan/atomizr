package collection

import (
	"fmt"
	"os"

	"go.yaml.in/yaml/v3"
)

// Fragrance is a single fragrance in the user's collection.
type Fragrance struct {
	Name          string   `yaml:"name"`
	Brand         string   `yaml:"brand"`
	Family        string   `yaml:"family"`
	Concentration string   `yaml:"concentration"`
	Intensity     string   `yaml:"intensity"`
	Seasons       []string `yaml:"seasons"`
	Occasions     []string `yaml:"occasions"`
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
	return fragrances, nil
}
