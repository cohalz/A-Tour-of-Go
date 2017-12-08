package main

import (
	"fmt"
	"sync"
)

type SafeIsVisitedMap struct {
	isVisited map[string]bool
	mux       sync.Mutex
}

func (m *SafeIsVisitedMap) Visit(url string) {
	m.mux.Lock()
	m.isVisited[url] = true
	m.mux.Unlock()
}

func (m *SafeIsVisitedMap) IsVisited(url string) bool {
	m.mux.Lock()
	defer m.mux.Unlock()
	return m.isVisited[url]
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
//クロール済みかどうかをmで判定，書込み
func Crawl(url string, depth int, fetcher Fetcher, m *SafeIsVisitedMap) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}

	//フェッチする前に判定
	if m.IsVisited(url) {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	//フェッチしたら訪問済みに
	m.Visit(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		Crawl(u, depth-1, fetcher, m)
	}
	return
}

func main() {
	//mapは自動で初期化されないため必ずmakeを呼ぶ
	//1つの実体を引き回したいのでポインタ
	m := &SafeIsVisitedMap{isVisited: make(map[string]bool)}
	Crawl("http://golang.org/", 4, fetcher, m)
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
