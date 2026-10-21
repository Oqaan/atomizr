package collection

// Intensity describes how strong a fragrance projects.
// The zero value means the intensity was not set.
type Intensity int

const (
	Light Intensity = iota + 1
	Moderate
	Strong
	Beast
)
