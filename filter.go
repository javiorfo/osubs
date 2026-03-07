package osubs

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/javiorfo/osubs/lang"
	"github.com/javiorfo/osubs/order"
	"github.com/javiorfo/steams/v2"
)

type filter struct {
	// Year to filter by (0 means no filter).
	year uint
	// Languages to filter by.
	languages []lang.Language
	// Sorting order.
	order order.By
}

func (f filter) create() string {
	var year string
	if f.year != 0 {
		year = strconv.Itoa(int(f.year))
	}

	langs := steams.FromSlice(f.languages).
		MapToString(func(l lang.Language) string { return l.Code() }).
		Collect()

	var orderBy string
	switch f.order {
	case order.Uploaded:
		orderBy = "/sort-5/asc-0"
	case order.Rating:
		orderBy = "/sort-6/asc-0"
	case order.Downloads:
		orderBy = "/sort-7/asc-0"
	}

	return fmt.Sprintf("&SubLanguageID=%s&MovieYearSign=1&MovieYear=%s%s", strings.Join(langs, ","), year, orderBy)
}
