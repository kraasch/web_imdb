package main

import (
  // read input.
  "os"
  // make web request to imdb.
  "github.com/StalkR/imdb"
	"github.com/StalkR/httpcache"
  "net/http"
  "time"
  // print output.
  "fmt"
)

const userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36"
const cacheTTL = 24 * time.Hour
var client *http.Client

type customTransport struct { // customTransport implements http.RoundTripper interface to add some headers.
	http.RoundTripper
}

func (e *customTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set("Accept-Language", "en") // avoid IP-based language detection
	r.Header.Set("User-Agent", userAgent)
	return e.RoundTripper.RoundTrip(r)
}

func AskImdb(title string) (string) {

  // init client.
	if _, err := os.Stat("cache"); err == nil {
		client, err = httpcache.NewPersistentClient("cache", cacheTTL)
		if err != nil {
			panic(err)
		}
	} else {
		client = httpcache.NewVolatileClient(cacheTTL, 1024)
	}
	client.Transport = &customTransport{client.Transport}

  // make web request.
  s := ""
  s += fmt.Sprintln("\n> Search title")
	results, err := imdb.SearchTitle(client, title)

  // evaluate and print results.
	if err != nil { // has error.
    s += fmt.Sprintf("imdb.SearchTitle(%s) error: %v", title, err)
	}
	if len(results) == 0 { // no results found.
    s += fmt.Sprintf("imdb.SearchTitle(%s) error: %d results found.", title, len(results))
	}
  // print results.
  s += fmt.Sprintf("%d titles found.\n", len(results))

  // get rating for title.
  if len(results) > 0 {
    fmt.Println()
    s += fmt.Sprintln("> Retrieve rating")
    var id = results[0].ID
    // make another request.
    titleResult, err := imdb.NewTitle(client, id)
    if err != nil {
      s += fmt.Sprintf("NewTitle(%s) error: %v\n", id, err)
    } else {
      s += fmt.Sprintf("First hit: '%s' (%d), duration '%s'.\n", titleResult.Name, titleResult.Year, titleResult.Duration)
      s += fmt.Sprintf("Rating: '%s'.\n", titleResult.Rating)
    }
  }
  return s
}

func main() {
  // read input.
  if len(os.Args) < 1 {
    fmt.Println("Usage: <program_name> <movie_name>")
    os.Exit(1)
  }
  title := os.Args[1]
  // make request.
  result := AskImdb(title)
  fmt.Println(result)
}

