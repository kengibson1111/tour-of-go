package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// SafeMap is safe to use concurrently.
type SafeMap struct {
	v   map[string]bool
	mux sync.Mutex
}

// Set acknowledges map entry for the given key.
func (m *SafeMap) Set(key string) {
	m.mux.Lock()
	// Lock so only one goroutine at a time can access the map m.v.
	m.v[key] = true
	m.mux.Unlock()
}

// Value returns the current value of the map for the given key.
func (m *SafeMap) Value(key string) (bool, bool) {
	m.mux.Lock()
	// Lock so only one goroutine at a time can access the map m.v.
	defer m.mux.Unlock()

	val, ok := m.v[key]
	return val, ok 
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, m SafeMap) {
	if depth <= 0 {
		return
	}

	_, ok := m.Value(url)
	if !ok {
		m.Set(url)
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("found: %s %q\n", url, body)
		for _, u := range urls {
			go Crawl(u, depth-1, fetcher, m)
		}
	}

	return
}

func main() {
	m := SafeMap{v: make(map[string]bool)}
	go Crawl("http://golang.org/", 4, fetcher, m)

	time.Sleep(time.Second)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}

