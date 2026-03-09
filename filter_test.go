package osubs

import (
	"testing"

	"github.com/javiorfo/osubs/lang"
	"github.com/javiorfo/osubs/order"
)

func TestFilter_Create(t *testing.T) {
	tests := []struct {
		name     string
		input    filter
		expected string
	}{
		{
			name: "Full filter with year, languages, and sorting",
			input: filter{
				year:      2023,
				languages: []lang.Language{lang.Spanish, lang.English},
				order:     order.Rating,
			},
			expected: "&SubLanguageID=spa,eng&MovieYearSign=1&MovieYear=2023/sort-6/asc-0",
		},
		{
			name: "No year filter",
			input: filter{
				year:      0,
				languages: []lang.Language{lang.English},
				order:     order.Uploaded,
			},
			expected: "&SubLanguageID=eng&MovieYearSign=1&MovieYear=/sort-5/asc-0",
		},
		{
			name: "Multiple languages, download sorting",
			input: filter{
				year:      1994,
				languages: []lang.Language{lang.SpanishLA, lang.English},
				order:     order.Downloads,
			},
			expected: "&SubLanguageID=spl,eng&MovieYearSign=1&MovieYear=1994/sort-7/asc-0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.create()
			if got != tt.expected {
				t.Errorf("filter.create() = %q, want %q", got, tt.expected)
			}
		})
	}
}
