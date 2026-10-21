package collection

import "fmt"

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
