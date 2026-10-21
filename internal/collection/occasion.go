package collection

import (
	"fmt"

	"go.yaml.in/yaml/v3"
)

// Occasion is a situation a fragrance is suited for.
// The zero value means the occasion was not set.
type Occasion int

const (
	Office Occasion = iota + 1
	Casual
	Evening
)

// ParseOccasion converts a value like "office" from collection.yaml
// into an Occasion.
func ParseOccasion(s string) (Occasion, error) {
	switch s {
	case "office":
		return Office, nil
	case "casual":
		return Casual, nil
	case "evening":
		return Evening, nil
	}
	return 0, fmt.Errorf("invalid occasion %q (want office, casual or evening)", s)
}

// String returns the name used in collection.yaml, e.g. "office".
func (o Occasion) String() string {
	switch o {
	case Office:
		return "office"
	case Casual:
		return "casual"
	case Evening:
		return "evening"
	}
	return "unknown"
}

// UnmarshalYAML lets the YAML library parse occasion values
// with ParseOccasion, so typos are reported as errors.
func (o *Occasion) UnmarshalYAML(value *yaml.Node) error {
	var s string
	err := value.Decode(&s)
	if err != nil {
		return err
	}
	parsed, err := ParseOccasion(s)
	if err != nil {
		return err
	}
	*o = parsed
	return nil
}
