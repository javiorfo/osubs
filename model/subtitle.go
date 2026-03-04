package model

import "github.com/javiorfo/nilo"

type Subtitle struct {
	// Unique identifier for the subtitle.
	id int
	// Movie title associated with the subtitle.
	movie string
	// Optional name or description of the subtitle.
	name nilo.Option[string]
	// Language code of the subtitle (e.g., "eng" for English).
	language string
	// CD or disc information (e.g., "CD1", "CD2").
	cd string
	// Upload date or timestamp.
	uploaded string
	// Number of times the subtitle has been downloaded.
	downloads int
	// User rating for the subtitle.
	rating float32
	// Optional uploader's username.
	uploader nilo.Option[string]
	// Direct download link for the subtitle file.
	downloadLink string
}

type Movie struct {
	/// Unique identifier for the movie.
	id int
	/// Movie title.
	name string
	/// URL to search for subtitles for this movie.
	subtitlesLink string
}
