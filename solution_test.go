package solution

import (
	"reflect"
	"sort"
	"testing"
	"time"
)

type mockHtmlParser struct {
	urls  []string
	edges [][]int
}

func (m *mockHtmlParser) GetUrls(url string) []string {
	time.Sleep(10 * time.Millisecond)

	// Find the index of the current URL
	var urlIdx int
	for i, u := range m.urls {
		if u == url {
			urlIdx = i
			break
		}
	}

	// Find all edges from this URL and get the connected URLs
	var result []string
	for _, edge := range m.edges {
		if edge[0] == urlIdx {
			result = append(result, m.urls[edge[1]])
		}
	}
	return result
}

func TestCrawl(t *testing.T) {
	tests := []struct {
		name     string   // test name
		urls     []string // list of all URLs
		edges    [][]int  // edges as [from_idx, to_idx]
		startUrl string   // start URL
		expected []string // expected result
	}{
		// 2. load all test cases
		{
			name: "case 1: all same domain",
			urls: []string{
				"http://news.yahoo.com/news/topics/",
				"http://news.yahoo.com/news",
			},
			edges: [][]int{
				{0, 1},
				{1, 0},
			},
			startUrl: "http://news.yahoo.com/news/topics/",
			expected: []string{
				"http://news.yahoo.com/news",
				"http://news.yahoo.com/news/topics/",
			},
		},
		{
			name: "case 2: mixed domains",
			urls: []string{
				"http://news.google.com",
				"http://news.google.com/sports",
				"http://news.google.com/weather",
				"http://news.yahoo.com",
			},
			edges: [][]int{
				{0, 3},
				{0, 1},
				{1, 2},
				{3, 0},
			},
			startUrl: "http://news.google.com",
			expected: []string{
				"http://news.google.com",
				"http://news.google.com/sports",
				"http://news.google.com/weather",
			},
		},
		{
			name: "case 3: dead end",
			urls: []string{
				"http://leetcode.com/problems",
			},
			edges:    [][]int{},
			startUrl: "http://leetcode.com/problems",
			expected: []string{
				"http://leetcode.com/problems",
			},
		},
	}

	for _, tt := range tests {
		// 3. use t.Run to create subtests
		t.Run(tt.name, func(t *testing.T) {
			parser := &mockHtmlParser{
				urls:  tt.urls,
				edges: tt.edges,
			}

			actualOutput := Crawl(tt.startUrl, parser)

			sort.Strings(actualOutput)
			sort.Strings(tt.expected)

			if !reflect.DeepEqual(actualOutput, tt.expected) {
				t.Errorf("expected %v, but got %v", tt.expected, actualOutput)
			}
		})
	}
}
