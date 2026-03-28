package solution

import (
	"net/url"
	"sync"
)

func SameHost(url1 string, url2 string) bool {

	parsed1, err1 := url.Parse(url1)
	parsed2, err2 := url.Parse(url2)
	if err1 != nil || err2 != nil {
		return false
	}
	return parsed1.Hostname() == parsed2.Hostname()
}

// Crawl retrieves all URLs reachable from startUrl that share the same host.

func Crawl(startUrl string, htmlParser HtmlParser) []string {

	visited := make(map[string]bool)

	var mu sync.Mutex
	var wg sync.WaitGroup
	var crawlNode func(url string)
	crawlNode = func(url string) {
		defer wg.Done()
		for _, nextUrl := range htmlParser.GetUrls(url) {
			mu.Lock()
			if !visited[nextUrl] && SameHost(url, nextUrl) {
				visited[nextUrl] = true
				wg.Add(1)
				go crawlNode(nextUrl)
			}
			mu.Unlock()
		}
	}
	visited[startUrl] = true
	wg.Add(1)
	go crawlNode(startUrl)
	wg.Wait()
	result := make([]string, 0, len(visited))
	for k := range visited {
		result = append(result, k)
	}
	return result
}
