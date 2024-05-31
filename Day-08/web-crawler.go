package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, cache *safeCache, ch chan<- string, wg *sync.WaitGroup) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:

	// fmt.Printf("Called Crawl() with %q\n\n", url)
	defer wg.Done()

	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		ch <- fmt.Sprintln(err)
		return
	}
	// fmt.Printf("found: %s %q\n", url, body)
	ch <- fmt.Sprintf("found: %s %q\n %v\n\n", url, body, urls)
	for _, u := range urls {
		if ok := cache.find(u); !ok {
			cache.insert(u)
			// fmt.Printf("calling Crawl() with %q\n\n", u)
			go Crawl(u, depth-1, fetcher, cache, ch, wg)
		}
	}
	return
}

func main() {
	cache := safeCache{m: make(map[string]bool)}
	ch := make(chan string)
	var wg sync.WaitGroup
	wg.Add(5)
	go func() {
		Crawl("https://golang.org/", 4, fetcher, &cache, ch, &wg)
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type safeCache struct {
	mu sync.Mutex
	m  map[string]bool
}

func (cache *safeCache) insert(s string) {
	cache.mu.Lock()
	cache.m[s] = true
	cache.mu.Unlock()
}

func (cache *safeCache) find(s string) bool {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	if _, ok := cache.m[s]; ok {
		return true
	} else {
		return false
	}
}

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
