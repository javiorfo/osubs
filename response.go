package osubs

import (
	"github.com/javiorfo/nilo"
	"github.com/javiorfo/steams/v2"
)

type Subtitle struct {
	// Unique identifier for the subtitle.
	ID uint
	// MovieTitle title associated with the subtitle.
	MovieTitle string
	// Optional description of the subtitle.
	Description nilo.Option[string]
	// Language code of the subtitle (e.g., "eng" for English).
	Language string
	// CD or disc information (e.g., "CD1", "CD2").
	Cd string
	// Upload date.
	Uploaded string
	// Number of times the subtitle has been downloaded.
	Downloads int
	// Format of subtitle file (e.g., "srt", "sub", "txt").
	Format string
	// User rating for the subtitle.
	Rating float32
	// Optional uploader's username.
	Uploader nilo.Option[string]
	// Direct download link for the subtitle file.
	DownloadLink string
}

type Movie struct {
	// Unique identifier for the movie.
	ID uint
	// Movie title.
	Name string
	// URL to search for subtitles for this movie.
	SubtitlesLink string
}

type Page struct {
	// The starting index of the current page.
	From int
	// The ending index of the current page.
	To int
	// The Total number of items available.
	Total int
}

type response interface {
	isResponse()
}

type Result[T any] struct {
	Page  Page
	Items steams.It[T]
}

func (Result[T]) isResponse() {}

func (r Result[T]) Next() nilo.Option[Result[T]] {
	return nilo.Nil[Result[T]]()
}
