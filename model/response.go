package model

import "github.com/javiorfo/nilo"

type Page struct {
	// The starting index of the current page.
	from int
	// The ending index of the current page.
	to int
	// The total number of items available.
	total int
}

func NewPage(from, to, total int) Page {
	return Page{from, to, total}
}

func (p Page) From() int {
	return p.from
}

func (p Page) To() int {
	return p.to
}

func (p Page) Total() int {
	return p.total
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
