package main

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"math/big"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"time"
)

const (
	charset     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	idLength    = 5
	maxAttempts = 10
)

// link is one shortened URL as it is written to the store file.
type link struct {
	ID        string    `json:"id"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}

// store holds every known link, keyed by short ID, plus the file it came from.
type store struct {
	path  string
	links map[string]link
}

// go run main.go -url=https://github.com/milwad-dev
// go run main.go -id=aB3xZ
// go run main.go -list
func main() {
	target := flag.String("url", "", "URL to shorten")
	id := flag.String("id", "", "Short ID to expand back into its URL")
	list := flag.Bool("list", false, "List every stored URL")
	path := flag.String("store", "urls.json", "File the shortened URLs are kept in")
	flag.Parse()

	if *target == "" && *id == "" && !*list {
		log.Fatal("Must provide -url, -id or -list")
	}

	s, err := loadStore(*path)
	if err != nil {
		log.Fatal(err)
	}

	switch {
	case *target != "":
		shortened, isNew, err := s.shorten(*target)
		if err != nil {
			log.Fatal(err)
		}

		if !isNew {
			fmt.Printf("Already shortened: %s\n", shortened.ID)
			return
		}

		if err := s.save(); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Shorter URL: %s\n", shortened.ID)
	case *id != "":
		shortened, ok := s.links[*id]
		if !ok {
			log.Fatalf("No URL stored for ID %q", *id)
		}

		fmt.Printf("The actual URL: %s\n", shortened.URL)
	default:
		for _, shortened := range s.sorted() {
			fmt.Printf("%s\t%s\n", shortened.ID, shortened.URL)
		}
	}
}

// loadStore reads the store file. A missing file is not an error: the first
// run simply starts with an empty store.
func loadStore(path string) (*store, error) {
	s := &store{path: path, links: make(map[string]link)}

	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return s, nil
	}
	if err != nil {
		return nil, fmt.Errorf("reading store: %w", err)
	}

	var links []link
	if err := json.Unmarshal(data, &links); err != nil {
		return nil, fmt.Errorf("parsing store %s: %w", path, err)
	}

	for _, shortened := range links {
		s.links[shortened.ID] = shortened
	}

	return s, nil
}

// shorten adds a URL to the store and returns its link. The bool reports
// whether the link is new; a URL that was already shortened keeps its old ID
// so the same input never hands out two different short IDs.
func (s *store) shorten(rawURL string) (link, bool, error) {
	normalized, err := normalizeURL(rawURL)
	if err != nil {
		return link{}, false, err
	}

	if existing, ok := s.findByURL(normalized); ok {
		return existing, false, nil
	}

	id, err := s.generateID()
	if err != nil {
		return link{}, false, err
	}

	shortened := link{ID: id, URL: normalized, CreatedAt: time.Now()}
	s.links[id] = shortened

	return shortened, true, nil
}

// findByURL looks a URL up the slow way. The store is keyed by ID because that
// is the lookup that has to be fast; scanning is fine for a local CLI.
func (s *store) findByURL(rawURL string) (link, bool) {
	for _, shortened := range s.links {
		if shortened.URL == rawURL {
			return shortened, true
		}
	}

	return link{}, false
}

// save writes the store back to disk. The file is written to a temporary name
// first and renamed into place, so an interrupted run cannot truncate the
// store and lose every link in it.
func (s *store) save() error {
	data, err := json.MarshalIndent(s.sorted(), "", "  ")
	if err != nil {
		return fmt.Errorf("encoding store: %w", err)
	}

	temp, err := os.CreateTemp(filepath.Dir(s.path), filepath.Base(s.path)+".tmp")
	if err != nil {
		return fmt.Errorf("creating temporary store: %w", err)
	}
	defer os.Remove(temp.Name())

	if _, err := temp.Write(data); err != nil {
		temp.Close()
		return fmt.Errorf("writing store: %w", err)
	}
	if err := temp.Close(); err != nil {
		return fmt.Errorf("writing store: %w", err)
	}

	if err := os.Rename(temp.Name(), s.path); err != nil {
		return fmt.Errorf("replacing store: %w", err)
	}

	return nil
}

// sorted returns the links oldest first, so both -list and the store file read
// in the order the URLs were shortened.
func (s *store) sorted() []link {
	links := make([]link, 0, len(s.links))
	for _, shortened := range s.links {
		links = append(links, shortened)
	}

	sort.Slice(links, func(i, j int) bool {
		if links[i].CreatedAt.Equal(links[j].CreatedAt) {
			return links[i].ID < links[j].ID
		}

		return links[i].CreatedAt.Before(links[j].CreatedAt)
	})

	return links
}

// normalizeURL validates a URL and returns it in a canonical form, so the same
// address typed two slightly different ways maps to one short ID.
func normalizeURL(rawURL string) (string, error) {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("invalid URL %q: %w", rawURL, err)
	}

	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return "", fmt.Errorf("invalid URL %q: only http and https are supported", rawURL)
	}
	if parsed.Host == "" {
		return "", fmt.Errorf("invalid URL %q: missing host", rawURL)
	}

	return parsed.String(), nil
}

// generateID returns a short ID that is not in the store yet.
func (s *store) generateID() (string, error) {
	for attempt := 0; attempt < maxAttempts; attempt++ {
		id, err := randomID()
		if err != nil {
			return "", err
		}

		if _, taken := s.links[id]; !taken {
			return id, nil
		}
	}

	return "", errors.New("could not find an unused ID, the store may be full")
}

// randomID builds one ID from crypto/rand, so IDs cannot be guessed by
// replaying the sequence a seeded generator would produce.
func randomID() (string, error) {
	id := make([]byte, idLength)

	for i := range id {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", fmt.Errorf("reading random data: %w", err)
		}

		id[i] = charset[index.Int64()]
	}

	return string(id), nil
}
