package collection

import (
	"fmt"

	"go.yaml.in/yaml/v3"
)

// Season is a time of year a fragrance is suited for.
// The zero value means the season was not set.
type Season int

const (
	Spring Season = iota + 1
	Summer
	Autumn
	Winter
)

// ParseSeason converts a value like "autumn" from collection.yaml
// into a Season.
func ParseSeason(s string) (Season, error) {
	switch s {
	case "spring":
		return Spring, nil
	case "summer":
		return Summer, nil
	case "autumn":
		return Autumn, nil
	case "winter":
		return Winter, nil
	}
	return 0, fmt.Errorf("invalid season %q (want spring, summer, autumn or winter)", s)
}

// String returns the name used in collection.yaml, e.g. "autumn".
func (s Season) String() string {
	switch s {
	case Spring:
		return "spring"
	case Summer:
		return "summer"
	case Autumn:
		return "autumn"
	case Winter:
		return "winter"
	}
	return "unknown"
}

// UnmarshalYAML lets the YAML library parse season values
// with ParseSeason, so typos are reported as errors.
func (s *Season) UnmarshalYAML(value *yaml.Node) error {
	var text string
	err := value.Decode(&text)
	if err != nil {
		return err
	}
	parsed, err := ParseSeason(text)
	if err != nil {
		return err
	}
	*s = parsed
	return nil
}
