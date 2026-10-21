package collection

import (
	"fmt"

	"go.yaml.in/yaml/v3"
)

// Intensity describes how strong a fragrance projects.
// The zero value means the intensity was not set.
type Intensity int

const (
	Light Intensity = iota + 1
	Moderate
	Strong
	Beast
)

// ParseIntensity converts a value like "moderate" from collection.yaml
// into an Intensity.
func ParseIntensity(s string) (Intensity, error) {
	switch s {
	case "light":
		return Light, nil
	case "moderate":
		return Moderate, nil
	case "strong":
		return Strong, nil
	case "beast":
		return Beast, nil
	}
	return 0, fmt.Errorf("invalid intensity %q (want light, moderate, strong or beast)", s)
}

// String returns the name used in collection.yaml, e.g. "moderate".
func (i Intensity) String() string {
	switch i {
	case Light:
		return "light"
	case Moderate:
		return "moderate"
	case Strong:
		return "strong"
	case Beast:
		return "beast"
	}
	return "unknown"
}

// UnmarshalYAML lets the YAML library parse intensity values
// with ParseIntensity, so typos are reported as errors.
func (i *Intensity) UnmarshalYAML(value *yaml.Node) error {
	var s string
	err := value.Decode(&s)
	if err != nil {
		return err
	}
	parsed, err := ParseIntensity(s)
	if err != nil {
		return err
	}
	*i = parsed
	return nil
}
