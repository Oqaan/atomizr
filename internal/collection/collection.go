package collection

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
