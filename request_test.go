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
	}

	switch v := resp.(type) {
	case Result[Subtitle]:
		for sub := range v.Items {
			t.Log(sub)
		}
	default:
		t.Fatalf("must be Result[Subtitle] %v", v)
	}
}

func TestSearchMovie(t *testing.T) {
	resp, err := Search("godfather", Language(lang.Spanish))
	if err != nil {
		t.Fatalf("calling %v", err)
	}

	if v, ok := resp.(Result[Movie]); !ok {
		t.Fatalf("must be Result[Movie] and got %v", v)
	}

	switch v := resp.(type) {
	case Result[Movie]:
		for movie := range v.Items {
			t.Log(movie)
		}
		t.Log(v.Page)
	default:
		t.Fatalf("must be Result[Movie] %v", v)
	}
}
