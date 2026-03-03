package main

import (
	"fmt"
	"regexp"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	c.OnHTML("div.content", func(e *colly.HTMLElement) {
		// si div#msg no existe es movies
		e.ForEach("div#msg", func(_ int, el *colly.HTMLElement) {
			// los primeros 3 son 1 - 2 of 2
			fmt.Println("page:", regexp.MustCompile(`\d+`).FindAllString(el.Text, -1))
		})

		e.ForEach("table#search_results", func(_ int, el *colly.HTMLElement) {
			fmt.Println("subs:", el.Text)
		})
	})

	c.OnError(func(r *colly.Response, err error) {
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://www.opensubtitles.org/en/search2?MovieName=godfather&SubLanguageID=spa&id=8&action=search")
}

