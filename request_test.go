package osubs

import "testing"

func TestSearch(t *testing.T) {
	resp, err := search("https://www.opensubtitles.org/en/search2?MovieName=holdovers&SubLanguageID=spa")
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
