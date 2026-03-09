package osubs

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/javiorfo/osubs/lang"
	"github.com/javiorfo/osubs/order"
	"github.com/javiorfo/steams/v2"
)

// filter defines the internal criteria for building a search query.
type filter struct {
	// year represents the release year. A value of 0 indicates no year filter.
	year uint
	// languages is a collection of Language interfaces to include in the search.
	languages []lang.Language
	// order determines the sorting priority (e.g., Uploaded, Rating).
	order order.By
}

// create assembles the final query string used for the external API request.
// It formats the languages, year, and sorting parameters into a URL-encoded string.
func (f filter) create() string {
	var year string
	if f.year != 0 {
		year = strconv.Itoa(int(f.year))
	}

	return fmt.Sprintf("&SubLanguageID=%s&MovieYearSign=1&MovieYear=%s%s", f.languagesToString(), year, f.orderToString())
}

// languagesToString converts the slice of Language interfaces into a
// comma-separated string of language codes (e.g., "en,fr,es").
func (f filter) languagesToString() string {
	langs := steams.FromSlice(f.languages).
		MapToString(func(l lang.Language) string { return l.Code() }).
		Collect()
	return strings.Join(langs, ",")
}

// orderToString maps the internal order.By type to the specific
// URL path segments required by the remote server.
func (f filter) orderToString() string {
	var orderBy string
	switch f.order {
	case order.Uploaded:
		orderBy = "/sort-5/asc-0"
	case order.Rating:
		orderBy = "/sort-6/asc-0"
	case order.Downloads:
		orderBy = "/sort-7/asc-0"
	}
	return orderBy
}
