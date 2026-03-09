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
		for _, sub := range v.Items {
			t.Log(sub)
		}
		v.Items[0].Download("/tmp")
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
		t.Log(v.Page)
		for i, movie := range v.Items {
			if i == 0 {
				subs, _ := movie.SearchSubtitles()
				t.Log(subs)
			}
			t.Log(movie)
		}
		/* 		t.Log(v.Next())
		   		t.Log(v.Page)
		   		t.Log(v.Items.Collect())
		   		t.Log(v.Next())
		   		t.Log(v.Page)
		   		t.Log(v.Items.Collect()) */

	default:
		t.Fatalf("must be Result[Movie] %v", v)
	}
}
