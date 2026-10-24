package collection

import "fmt"

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
