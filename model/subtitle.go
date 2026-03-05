package model

import "github.com/javiorfo/nilo"

type Subtitle struct {
	// Unique identifier for the subtitle.
	ID int
	// Movie title associated with the subtitle.
	Movie string
	// Optional name or description of the subtitle.
	Name nilo.Option[string]
	// Language code of the subtitle (e.g., "eng" for English).
	Language string
	// CD or disc information (e.g., "CD1", "CD2").
	Cd string
	// Upload date or timestamp.
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
	/// Unique identifier for the movie.
	ID int
	/// Movie title.
	Name string
	/// URL to search for subtitles for this movie.
	SubtitlesLink string
}
