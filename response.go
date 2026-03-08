package osubs

import (
	"errors"
	"fmt"
	"iter"
	"regexp"
	"strconv"
	"strings"

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

// Zipped or unzipped
func (s Subtitle) Download(path string) error {
	return nil
}

type Movie struct {
	// Unique identifier for the movie.
	ID uint
	// Movie title.
	Name string
	f    filter
}

func (m Movie) SearchSubtitles() (Result[Subtitle], error) {
	subtitlesLink := fmt.Sprintf(
		"https://www.opensubtitles.org/en/search/sublanguageid-%s/idmovie-%d%s",
		m.f.languagesToString(),
		m.ID,
		m.f.orderToString(),
	)

	resp, err := search(subtitlesLink, m.f)
	return resp.(Result[Subtitle]), err
}

type Page struct {
	// The starting index of the current page.
	From int
	// The ending index of the current page.
	To int
	// The Total number of items available.
	Total int
}

func newPage(text string) (*Page, error) {
	numbers := regexp.MustCompile(`\d+`).FindAllString(text, -1)
	if len(numbers) < 3 {
		return nil, errors.New("could not set Page values (less than 3)")
	}

	from, err := strconv.Atoi(numbers[0])
	if err != nil {
		return nil, err
	}
	to, err := strconv.Atoi(numbers[1])
	if err != nil {
		return nil, err
	}
	total, err := strconv.Atoi(numbers[2])
	if err != nil {
		return nil, err
	}

	return &Page{From: from, To: to, Total: total}, nil
}

type response interface {
	isResponse()
}

type Result[T any] struct {
	Page  Page
	Items steams.It[T]
	url   string
	f     filter
}

func newResult[T any](url string) Result[T] {
	if before, _, ok := strings.Cut(url, "/offset"); ok {
		return Result[T]{url: before}
	}
	return Result[T]{url: url}
}

func (Result[T]) isResponse() {}

func (r *Result[T]) Next() (bool, error) {
	if r.Page.Total > r.Page.From+40 {
		r.Page.From += 39
		resp, err := search(fmt.Sprintf("%s/offset-%d", r.url, r.Page.From), r.f)
		if err != nil {
			return true, err
		}

		*r = resp.(Result[T])
		return true, nil
	}
	return false, nil
}

func (r *Result[T]) Iter() iter.Seq2[*Result[T], error] {
	return func(yield func(*Result[T], error) bool) {
		if !yield(r, nil) {
			return
		}

		for {
			next, err := r.Next()
			if err != nil {
				yield(nil, err)
				return
			}
			if !next {
				break
			}

			if !yield(r, nil) {
				return
			}
		}
	}
}
