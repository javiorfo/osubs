package osubs

import (
	"testing"

	"github.com/javiorfo/osubs/lang"
)

func TestSearchSubtitleRedirect(t *testing.T) {
	resp, err := Search("holdovers", Language(lang.Spanish, lang.SpanishLA))
	if err != nil {
		t.Fatalf("calling %v", err)
	}

	if v, ok := resp.(Result[Subtitle]); !ok {
		t.Fatalf("must be Result[Subtitle] and got %v", v)
	} else {
		v.Items.First().Consume(func(s Subtitle) {
			if err := s.Download("/tmp"); err != nil {
				t.Fatalf("downloading subtitle %v", err)
			}
		})
	}
}

func TestSearchMovie(t *testing.T) {
	resp, err := Search("godfather", Language(lang.Spanish))
	if err != nil {
		t.Fatalf("calling %v", err)
	}

	switch v := resp.(type) {
	case Result[Movie]:
		t.Logf("Initial page: %v\n", v.Page)
		v.Items.First().Consume(func(m Movie) {
			subs, err := m.SearchSubtitles()
			if err != nil {
				t.Fatalf("searching subtitle from movie %v", err)
			}
			t.Logf("Subtitles inside movie: %v\n", subs)
		})

		t.Log(v.Next())
		t.Logf("Next page: %v\n", v.Page)
		t.Log(v.Items.Collect())

		t.Log(v.Next())
		t.Logf("Last page: %v\n", v.Page)
		t.Log(v.Items.Collect())
	default:
		t.Fatalf("must be Result[Movie] %v", v)
	}
}
