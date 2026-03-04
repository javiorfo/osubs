package request

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/gocolly/colly/v2"
	"github.com/javiorfo/osubs/model"
)

func Search() (resp model.Response, err error) {
	c := colly.NewCollector()

	c.OnHTML("div.content", func(e *colly.HTMLElement) {
		// Checking for subtitle pagination
		if e.DOM.Find("div#msg").Length() > 0 {
			sr := model.SubtitleResponse{}

			e.ForEach("div#msg", func(_ int, el *colly.HTMLElement) {
				numbers := regexp.MustCompile(`\d+`).FindAllString(el.Text, -1)
				if len(numbers) < 3 {
					return
				}

				from, err1 := strconv.Atoi(numbers[0])
				if err != nil {
					err = err1
					return
				}
				to, err2 := strconv.Atoi(numbers[1])
				if err != nil {
					err = err2
					return
				}
				total, err3 := strconv.Atoi(numbers[2])
				if err != nil {
					err = err3
					return
				}

				sr.Page = model.NewPage(from, to, total)
			})

			e.ForEach("table#search_results", func(_ int, el *colly.HTMLElement) {
				fmt.Println("subs:", el.Text)
			})

			resp = sr
		} else {
			mr := model.MovieResponse{}
			e.ForEach("table#search_results", func(_ int, el *colly.HTMLElement) {
				fmt.Println("subs:", el.Text)
			})
			resp = mr
		}
	})

	c.OnError(func(r *colly.Response, e error) {
		err = e
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	err = c.Visit("https://www.opensubtitles.org/en/search2?MovieName=holdovers&SubLanguageID=spa&id=8&action=search")

	return
}
