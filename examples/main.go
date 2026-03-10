package main

import (
	"log"

	"github.com/javiorfo/osubs"
	"github.com/javiorfo/osubs/lang"
	"github.com/javiorfo/osubs/order"
)

func main() {
	resp, err := osubs.Search("godfather", osubs.Language(lang.Spanish, lang.SpanishLA), osubs.Year(1972))
	if err != nil {
		log.Fatal(err)
	}
	processResponse(resp)

	resp, err = osubs.Search("pulp fiction", osubs.Language(lang.French, lang.German))
	if err != nil {
		log.Fatal(err)
	}
	processResponse(resp)

	resp, err = osubs.Search("terminator", osubs.Order(order.Downloads))
	if err != nil {
		log.Fatal(err)
	}
	processResponse(resp)
}

func processResponse(r osubs.Response) {
	switch s := r.(type) {
	case osubs.Result[osubs.Subtitle]:
		for sub := range s.Items {
			log.Printf("%+v\n", sub)
		}

	case osubs.Result[osubs.Movie]:
		for movie := range s.Items {
			log.Printf("%+v\n", movie)
		}

		s.Items.First().Consume(func(m osubs.Movie) {
			resp, err := m.SearchSubtitles()
			if err != nil {
				log.Fatal(err)
			}
			processResponse(resp)
		})

	default:
		log.Println("no results")
	}
}
