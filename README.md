# osubs
*Library to search subtitles from opensubtitles.org*

## Caveats
- This library requires Go 1.26
- This library provides a high-level, ergonomic API for searching and retrieving subtitles (only from movies) and related metadata from `opensubtitles.org`.

## Insstallation
```bash
go get -u github.com/javiorfo/osubs@latest
```

## Example
```go
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
```

## Details
- Searching subtitles from `opensubtitles.org` could return a list of movies or a list of subtitles of the movie searched (if the text and filter are more exactly). For that matter Response is a sealed interface


---

### Donate
- **Bitcoin** [(QR)](https://raw.githubusercontent.com/javiorfo/img/master/crypto/bitcoin.png)  `1GqdJ63RDPE4eJKujHi166FAyigvHu5R7v`
- [Paypal](https://www.paypal.com/donate/?hosted_button_id=FA7SGLSCT2H8G)
