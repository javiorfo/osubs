package request

import (
	"testing"

	"github.com/javiorfo/osubs/model"
)

func TestSearch(t *testing.T) {
	resp, err := Search()
	if err != nil {
		t.Fatalf("calling %v", err)
	}

	switch v := resp.(type) {
	case model.SubtitleResponse:
		t.Log(v)
	default:
		t.Fatalf("must be SubtitleResponse %v", v)
	}
}
