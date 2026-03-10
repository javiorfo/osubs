package osubs

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/javiorfo/nilo"
	"github.com/javiorfo/steams/v2"
)

// Subtitle represents the metadata and download information for a specific subtitle file.
type Subtitle struct {
	// ID is the unique identifier for the subtitle on the remote provider.
	ID uint
	// MovieTitle is the name of the film or show associated with this subtitle and the year of release.
	MovieTitle string
	// Description provides optional context or version details (e.g., "Godfather.1080p.DVD").
	Description nilo.Option[string]
	// Language is the subtitle language.
	Language string
	// Cd indicates disc information for multi-part releases (e.g., "CD1").
	Cd string
	// Uploaded is the date when the subtitle file was first provided.
	Uploaded string
	// Downloads tracks the popularity of this specific subtitle entry.
	Downloads int
	// Format is the file extension (e.g., "srt", "sub").
	Format string
	// Rating represents the community-assigned quality score.
	Rating float32
	// Uploader is the username of the community member who provided the file.
	Uploader nilo.Option[string]
	// DownloadLink is the direct URL used to fetch the subtitle archive.
	DownloadLink string
}

// Download fetches the subtitle archive from the remote server, extracts the
// relevant file matching the Subtitle's format, and saves it to the specified path.
//
// The resulting file is named using the MovieTitle and Format fields.
// The download link points to a ZIP archive, this method automatically
// iterates through the archive to find and extract the correct file.
func (s Subtitle) Download(path string) error {
	resp, err := http.Get(s.DownloadLink)
	if err != nil {
		return fmt.Errorf("failed to fetch subtitle from %s: %w", s.DownloadLink, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad server response: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	zipReader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		return fmt.Errorf("failed to open zip archive: %w", err)
	}

	for _, file := range zipReader.File {
		if !strings.HasSuffix(strings.ToLower(file.Name), s.Format) || file.FileInfo().IsDir() {
			continue
		}

		zippedFile, err := file.Open()
		if err != nil {
			return err
		}
		defer zippedFile.Close()

		filePath := filepath.Join(path, fmt.Sprintf("%s.%s", s.MovieTitle, s.Format))
		destFile, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("could not create file %s: %w", filePath, err)
		}
		defer destFile.Close()

		_, err = io.Copy(destFile, zippedFile)
		if err != nil {
			return fmt.Errorf("write error: %w", err)
		}
	}
	return nil
}

// Movie represents a cinematic title and its associated metadata.
// It contains a filter to refine subtitle searches specifically for this title.
type Movie struct {
	// ID is the unique identifier for the movie on the provider's database.
	ID uint
	// Name is the official title of the movie and year of release.
	Name string
	// f is the internal filter configuration used when searching for subtitles.
	f filter
}

// SearchSubtitles performs a remote request to find subtitles matching the movie's
// ID and the configured filter criteria (language, year, and sort order).
//
// It returns a Result containing a slice of Subtitle objects or an error
// if the network request or parsing fails.
func (m Movie) SearchSubtitles() (Result[Subtitle], error) {
	subtitlesLink := fmt.Sprintf(
		"https://www.opensubtitles.org/en/search/sublanguageid-%s/idmovie-%d%s",
		m.f.languagesToString(),
		m.ID,
		m.f.orderToString(),
	)
	fmt.Println(subtitlesLink)

	resp, err := search(subtitlesLink, m.f)
	return resp.(Result[Subtitle]), err
}

// Page represents the pagination info.
type Page struct {
	// The starting index of the current page.
	From int
	// The ending index of the current page.
	To int
	// The Total number of items available.
	Total int
}

// newPage parses a pagination summary string to create a Page reference.
// It expects a text input containing at least three integers representing
// the "from", "to", and "total" values (e.g., "Displaying 1 to 40 of 70").
//
// If the input text contains fewer than three numeric sequences, or if
// the found sequences cannot be parsed as integers, it returns an error.
func newPage(text string) (*Page, error) {
	numbers := regexp.MustCompile(`\d+`).FindAllString(text, -1)
	if len(numbers) < 3 {
		return nil, errors.New("could not set Page values (less than 3)")
	}

	from, err := strconv.Atoi(numbers[0])
	if err != nil {
		return nil, fmt.Errorf("getting 'from' Page: %w", err)
	}
	to, err := strconv.Atoi(numbers[1])
	if err != nil {
		return nil, fmt.Errorf("getting 'to' page: %w", err)
	}
	total, err := strconv.Atoi(numbers[2])
	if err != nil {
		return nil, fmt.Errorf("getting 'total' page: %w", err)
	}

	return &Page{From: from, To: to, Total: total}, nil
}

// Response is a sealed interface used to categorize internal search results.
type Response interface {
	isResponse()
}

// Result holds a paginated collection of items of type T.
// It maintains the internal state required to fetch subsequent pages from the source.
type Result[T any] struct {
	// Page contains metadata about the current position and total results.
	Page Page
	// Items is stream iterator of data retrieved for the current page.
	Items steams.It[T]
	// url is the base endpoint used for pagination.
	url string
	// f is the filter configuration applied to the search.
	f filter
}

// newResult creates a Result instance, stripping any existing offset
// from the URL to establish a clean base for pagination.
func newResult[T any](url string) Result[T] {
	if before, _, ok := strings.Cut(url, "/offset"); ok {
		return Result[T]{url: before}
	}
	return Result[T]{url: url}
}

func (Result[T]) isResponse() {}

// OpenSubtitles typically uses 40-item increments for pagination.
const pageIncrement = 40

// Next advances the Result to the next page of items.
// It updates the receiver in-place with the new data.
//
// It returns true if a new page was successfully fetched, and false if
// no more items are available. An error is returned if the network
// request or parsing fails.
func (r *Result[T]) Next() (bool, error) {
	if r.Page.Total > r.Page.From+pageIncrement {
		if r.Page.From == 1 {
			r.Page.From -= 1
		}
		r.Page.From += pageIncrement

		resp, err := search(fmt.Sprintf("%s/offset-%d", r.url, r.Page.From), r.f)
		if err != nil {
			return true, err
		}

		*r = resp.(Result[T])
		return true, nil
	}
	return false, nil
}

// Iter returns a stream-compatible iterator (steams.It2) that yields the
// current Result and all subsequent pages.
//
// This allows for clean iteration over all available search results
// using a range-like functional approach.
func (r *Result[T]) Iter() steams.It2[*Result[T], error] {
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
