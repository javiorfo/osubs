package osubs

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
	"github.com/javiorfo/nilo"
	"github.com/javiorfo/osubs/lang"
	"github.com/javiorfo/osubs/order"
	"github.com/javiorfo/steams/v2"
)

func Search(movieName string, filters ...filterOpts) (response, error) {
	if movieName == "" {
		return nil, errors.New("movie name must not be empty")
	}

	f := &filter{}
	for _, opt := range filters {
		opt(f)
	}

	return search(fmt.Sprintf("https://www.opensubtitles.org/en/search2?MovieName=%s%s", movieName, f.create()))
}

type filterOpts func(*filter)

func Year(y uint) filterOpts {
	return func(f *filter) {
		f.year = y
	}
}

func Language(l ...lang.Language) filterOpts {
	return func(f *filter) {
		if len(l) > 0 {
			f.languages = l
		}
	}
}

func Order(o order.By) filterOpts {
	return func(f *filter) {
		switch f.order {
		case order.Uploaded, order.Downloads, order.Rating:
			f.order = o
		default:
			f.order = order.Uploaded
		}
	}
}

func search(url string) (resp response, searchErr error) {
	c := colly.NewCollector()
	extensions.RandomUserAgent(c)

	c.OnHTML("div.content", func(e *colly.HTMLElement) {
		// Checking for subtitle pagination
		if e.DOM.Find("div#msg").Length() > 0 {
			sr := Result[Subtitle]{}

			e.ForEach("div#msg", func(_ int, el *colly.HTMLElement) {
				numbers := regexp.MustCompile(`\d+`).FindAllString(el.Text, -1)
				if len(numbers) < 3 {
					searchErr = errors.New("could not set Page values (less than 3)")
					return
				}

				from, err := strconv.Atoi(numbers[0])
				if err != nil {
					searchErr = err
					return
				}
				to, err := strconv.Atoi(numbers[1])
				if err != nil {
					searchErr = err
					return
				}
				total, err := strconv.Atoi(numbers[2])
				if err != nil {
					searchErr = err
					return
				}

				sr.Page = Page{From: from, To: to, Total: total}
			})

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
						searchErr = err
						return
					}

					sub.ID = uint(id)
					sub.DownloadLink = fmt.Sprintf("https://dl.opensubtitles.org/en/download/sub/%d", id)

					row.ForEach("td", func(i int, el *colly.HTMLElement) {
						item := strings.TrimSpace(el.Text)
						switch i {
						case 0:
							names := steams.FromSlice(strings.Split(item, "\n"))
							sub.MovieTitle = names.First().MapOrDefault(func(s string) string {
								return strings.TrimSpace(s)
							})
							sub.Description = names.Nth(1).AndThen(func(s string) nilo.Option[string] {
								trimmed := strings.TrimSpace(s)
								if trimmed == "" {
									return nilo.Nil[string]()
								}
								cleanStr, _, _ := strings.Cut(trimmed, "Watch")
								if _, after, ok := strings.Cut(cleanStr, ")"); ok {
									cleanStr = strings.TrimSpace(after)
								}
								return nilo.Value(cleanStr)
							})
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

			sr.Items = steams.FromSlice(subtitles)
			resp = sr
		} else if e.DOM.Find("div.msg none").Length() > 0 {
			mr := Result[Movie]{}
			e.ForEach("table#search_results tr", func(i int, el *colly.HTMLElement) {
				fmt.Println("subs:", el.Text)
			})
			resp = mr
		}
	})

	c.OnError(func(r *colly.Response, e error) {
		searchErr = e
	})

	searchErr = c.Visit(url)

	return
}
