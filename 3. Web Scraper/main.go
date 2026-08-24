package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

const appPort = ":8080"

var tmpl = template.Must(
	template.ParseFiles("index.html"),
)

type Data struct {
	URL             string
	Title           string
	Name            string
	Bio             string
	Location        string
	Followers       int
	Following       int
	RepositoryCount int
	Avatar          string
}

func main() {
	http.HandleFunc("/", handleIndex)

	fmt.Printf("Application running on http://localhost%s/\n", appPort)
	if err := http.ListenAndServe(appPort, nil); err != nil {
		log.Fatal(err)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	var data Data
	if url := r.URL.Query().Get("url"); url != "" {
		var err error
		if data, err = collectDataFromURL(url); err != nil {
			http.Error(w, "Scrape failed: "+err.Error(), http.StatusBadGateway)
			return
		}
	}

	var page bytes.Buffer
	if err := tmpl.Execute(&page, data); err != nil {
		log.Println("template error:", err)
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	page.WriteTo(w)
}

func collectDataFromURL(url string) (Data, error) {
	data := Data{URL: url}

	collector := colly.NewCollector()

	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	collector.OnError(func(r *colly.Response, err error) {
		fmt.Println("Blimey, an error occurred!:", err)
	})

	collector.OnHTML("html", func(e *colly.HTMLElement) {
		data.Title = strings.TrimSpace(e.DOM.Find("title").First().Text())
		data.Name = strings.TrimSpace(e.DOM.Find("span.p-name").First().Text())
		data.Bio = strings.TrimSpace(e.DOM.Find("div.p-note").First().Text())
		data.Location = strings.TrimSpace(e.DOM.Find("span.p-label").First().Text())
		data.Followers = parseCount(e.DOM.Find("a[href*='tab=followers']").First().Text())
		data.Following = parseCount(e.DOM.Find("a[href*='tab=following']").First().Text())
		data.RepositoryCount = parseCount(e.DOM.Find("a[href*='tab=repositories'] span.Counter").First().Text())

		if src := e.ChildAttr("img.avatar-user", "src"); src != "" {
			data.Avatar = e.Request.AbsoluteURL(src)
		}
	})

	if err := collector.Visit(url); err != nil {
		return Data{}, err
	}

	return data, nil
}

// parseCount pulls the first number out of text like "412 followers", coping
// with the "1.2k" shorthand GitHub uses for larger counts.
func parseCount(text string) int {
	for field := range strings.FieldsSeq(strings.ReplaceAll(text, ",", "")) {
		if field[0] < '0' || field[0] > '9' {
			continue
		}

		multiplier := 1.0
		if strings.HasSuffix(strings.ToLower(field), "k") {
			field = field[:len(field)-1]
			multiplier = 1000
		}

		value, err := strconv.ParseFloat(field, 64)
		if err != nil {
			continue
		}

		return int(value * multiplier)
	}

	return 0
}
