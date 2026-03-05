package model

import "github.com/javiorfo/nilo"

type Page struct {
	// The starting index of the current page.
	From int
	// The ending index of the current page.
	To int
	// The Total number of items available.
	Total int
}

type Response interface {
	isResponse()
}

type SubtitleResponse struct {
	Page      Page
	Subtitles []Subtitle
}

func (sr SubtitleResponse) isResponse() {}

func (sr SubtitleResponse) Next() nilo.Option[SubtitleResponse] {
	return nilo.Nil[SubtitleResponse]()
}

type MovieResponse struct {
	page   Page
	movies []Movie
}

func (mr MovieResponse) isResponse() {}

func (mr MovieResponse) Next() nilo.Option[MovieResponse] {
	return nilo.Nil[MovieResponse]()
}
