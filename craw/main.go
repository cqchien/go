package main

import (
	"fmt"
	"sync"
)

// In this exercise you'll use Go's concurrency features to parallelize a web crawler.

// Modify the Crawl function to fetch URLs in parallel without fetching the same URL twice.

// Hint: you can keep a cache of the URLs that have been fetched on a map, but maps alone are not safe for concurrent use!


type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type SafeMap struct {
	mu sync.Mutex
	v  map[string]bool
}

func (sm *SafeMap) Add(url string) bool {
	sm.mu.Lock()
	if sm.v[url] {
		sm.mu.Unlock()
		return false
	}
	sm.v[url] = true
	sm.mu.Unlock()
	return true
}

func Crawl(url string, depth int, fetcher Fetcher, sm *SafeMap, wg *sync.WaitGroup) {
	defer wg.Done()
	if depth <= 0 || !sm.Add(url) {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		wg.Add(1)
		go Crawl(u, depth-1, fetcher, sm, wg)
	}
}

func main() {
	var wg sync.WaitGroup
	sm := &SafeMap{v: make(map[string]bool)}
	wg.Add(1)
	Crawl("https://golang.org/", 4, fetcher, sm, &wg)
	wg.Wait()
}

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

var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
