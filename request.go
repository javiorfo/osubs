package osubs

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
	"github.com/javiorfo/nilo"
	"github.com/javiorfo/osubs/lang"
	"github.com/javiorfo/osubs/order"
	"github.com/javiorfo/steams/v2"
)

// filterOpts defines a function type used to configure a filter.
type filterOpts func(*filter)

// Year sets the release year criterion for the filter.
// A value of 0 typically disables the year filter.
func Year(y uint) filterOpts {
	return func(f *filter) {
		f.year = y
	}
}

// Language sets one or more languages to filter by.
// If no languages are provided, the filter remains unchanged.
func Language(l ...lang.Language) filterOpts {
	return func(f *filter) {
		if len(l) > 0 {
			f.languages = l
		}
	}
}

// Order sets the sorting preference for the search results.
// Valid options include order.Uploaded, order.Downloads, and order.Rating.
// If an invalid order is detected, it defaults to order.Uploaded.
func Order(o order.By) filterOpts {
	return func(f *filter) {
		switch o {
		case order.Uploaded, order.Downloads, order.Rating:
			f.order = o
		default:
			f.order = order.Uploaded
		}
	}
}

// Search initiates a subtitle or movie search on OpenSubtitles.
// It takes a movie title and optional filter configurations (Year, Language, Order).
//
// If the movieName is empty, it returns an error. It constructs the appropriate
// search URL and invokes the internal scraping engine.
func Search(movieName string, filters ...filterOpts) (response, error) {
	if movieName == "" {
		return nil, errors.New("movie name must not be empty")
	}

	f := &filter{}
	for _, opt := range filters {
		opt(f)
	}

	return search(fmt.Sprintf("https://www.opensubtitles.org/en/search2?MovieName=%s%s", movieName, f.create()), *f)
}

// search is the internal scraper that parses the OpenSubtitles HTML response.
// It uses the Colly framework to navigate the DOM and can return either
// a Result[Subtitle] or a Result[Movie] depending on whether the search
// led directly to a list of subtitles or a list of possible movie matches.
func search(url string, f filter) (resp response, searchErr error) {
	c := colly.NewCollector()
	extensions.RandomUserAgent(c)

	// The logic inside this callback handles two distinct page layouts:
	// 1. A table of Subtitles (found when the search is specific).
	// 2. A table of Movies (found when the search is ambiguous).
	c.OnHTML("div.content", func(e *colly.HTMLElement) {
		// Checking for subtitle pagination
		if e.DOM.Find("div#msg").Length() > 0 {
			sr := newResult[Subtitle](url)

			page, err := newPage(e.ChildText("div#msg"))
			if err != nil {
				searchErr = err
				return
			}
			sr.Page = *page

			var subtitles []Subtitle
			e.ForEach("table#search_results tr", func(i int, row *colly.HTMLElement) {
				if i == 0 {
					return
				}

				nameID := strings.TrimSpace(row.Attr("id"))
				if !strings.Contains(nameID, "ihtr") {
					sub := Subtitle{}

					id, err := strconv.Atoi(strings.TrimPrefix(nameID, "name"))
					if err != nil {
						searchErr = fmt.Errorf("converting id movie to number: %w", err)
						return
					}

					sub.ID = uint(id)
					sub.DownloadLink = fmt.Sprintf("https://dl.opensubtitles.org/en/download/sub/%d", id)

					row.ForEach("td", func(i int, el *colly.HTMLElement) {
						item := strings.TrimSpace(el.Text)
						switch i {
						case 0:
							sub.MovieTitle = formatMovieName(el.ChildText("[id] strong"))
							sub.Description = steams.FromSlice(strings.Split(item, "\n")).
								Nth(1).
								AndThen(func(s string) nilo.Option[string] {
									cleanStr, _, _ := strings.Cut(s, "Watch")
									if _, after, ok := strings.Cut(cleanStr, ")"); ok {
										cleanStr = strings.TrimSpace(after)
									}

									trimmed := strings.TrimSpace(cleanStr)
									if trimmed == "" {
										return nilo.Nil[string]()
									}
									return nilo.Value(trimmed)
								})
						case 1:
							sub.Language = el.ChildAttr("a", "title")
						case 2:
							sub.Cd = item
						case 3:
							sub.Uploaded = item[:8]
						case 4:
							downloads := steams.FromSlice(strings.Split(item, "\n"))
							sub.Downloads = downloads.First().MapToInt(func(s string) int {
								return nilo.Cast[int](strings.ReplaceAll(s, "x", "")).OrDefault()
							}).OrDefault()
							sub.Format = downloads.Nth(1).MapOrDefault(func(s string) string {
								return strings.TrimSpace(s)
							})
						case 5:
							sub.Rating = nilo.Cast[float32](item).OrDefault()
						case 8:
							if item != "" {
								sub.Uploader = nilo.Value(item)
							}
						default:
						}
					})

					subtitles = append(subtitles, sub)
				}
			})

			sr.Items = subtitles
			resp = sr
		} else if e.DOM.Find("div.msg.none").Length() > 0 {
			mr := newResult[Movie](url)

			page, err := newPage(e.ChildText("div.msg.none"))
			if err != nil {
				searchErr = err
				return
			}
			mr.Page = *page

			var movies []Movie

			e.ForEach("table#search_results tr", func(i int, row *colly.HTMLElement) {
				if i == 0 {
					return
				}

				movie := Movie{f: f}
				nameID := strings.TrimSpace(row.Attr("id"))
				id, err := strconv.Atoi(strings.TrimPrefix(nameID, "name"))
				if err != nil {
					searchErr = fmt.Errorf("converting id movie to number: %w", err)
					return
				}

				movie.ID = uint(id)
				movie.Name = formatMovieName(row.ChildText("td[id] strong"))

				movies = append(movies, movie)
			})

			mr.Items = movies
			resp = mr
		}
	})

	c.OnError(func(r *colly.Response, e error) {
		searchErr = fmt.Errorf("on request: %w, response: %v", e, r)
	})

	searchErr = c.Visit(url)

	return

}

// formatMovieName cleans up the raw title string extracted from the HTML.
// It merges multi-line title strings (e.g., Title + Year) into a single line.
func formatMovieName(raw string) string {
	movieName := strings.Split(strings.TrimSpace(raw), "\n")
	if len(movieName) > 1 {
		return fmt.Sprintf("%s %s", movieName[0], strings.TrimSpace(movieName[1]))
	}
	return ""
}
