package order

type By uint

const (
	// Sort by upload date (default).
	Uploaded By = iota
	// Sort by number of downloads.
	Downloads
	// Sort by rating.
	Rating
)
