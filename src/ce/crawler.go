package main

import (
	"fmt"
	"sync"
)

/*

https://tour.golang.org/concurrency/10

Exercise: Web Crawler
In this exercise you'll use Go's concurrency features to parallelize a web crawler.

Modify the Crawl function to fetch URLs in parallel without fetching the same URL twice.

Hint: you can keep a cache of the URLs that have been fetched on a map,
but maps alone are not safe for concurrent use!

 */

// *************************************
// Set implementation with synchronized methods
// *************************************

type IntSet struct {
	set map[string]bool
	mux sync.Mutex
}

func NewIntSet() *IntSet {
	return &IntSet{set: make(map[string]bool)}
}

func (set *IntSet) Add(i string) bool {
	set.mux.Lock()
	_, found := set.set[i]
	set.set[i] = true
	defer set.mux.Unlock()
	return !found	//False if it existed already
}

func (set *IntSet) Contains(i string) bool {
	set.mux.Lock()
	_, found := set.set[i]
	defer set.mux.Unlock()
	return found	//true if it existed already
}

func (set *IntSet) Remove(i string) {
	set.mux.Lock()
	delete(set.set, i)
	set.mux.Unlock()
}

// *************************************
// Fetcher interface
// *************************************

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// we the defet statement at the start because when
	// depth is zero it returns and before that 'defer'
	// is not set; we go into deadlock
	defer wg.Done()
	if depth <= 0 {
		return
	}
	fu.Add(url)
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		// we could not fetch, add back
		fu.Remove(url)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		//fmt.Printf("u: %s\n", u)
		if !fu.Contains(u) {
			wg.Add(1)
			//fmt.Printf("Spawning go routine to crawl %s \n", u);
			go Crawl(u, depth-1, fetcher)
		}
		// else it is present, no need to start another go routine to fetch
	}
	return
}

// *************************************
// Variables accessed by all go routines
// *************************************

// set to store url already fetched
var fu *IntSet
// wait group to tracked all spawned go routines
var wg sync.WaitGroup

// *************************************
// main for testing
// *************************************

func main() {
	// create a set
	fu = NewIntSet()
	// start the first go routine for the initial url
	wg.Add(1)
	go Crawl("https://golang.org/", 4, fetcher)
	wg.Wait()
}

// *************************************
// Fetcher implementation for testing
// *************************************

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

