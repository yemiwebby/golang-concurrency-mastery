package main

import (
	"strings"
	"sync"
	"testing"

	"golang.org/x/net/html"
)


func TestParseURL(t *testing.T) {
	urlString := "https://example.com/path"
	domain, scheme, err := parseURL(urlString)
	if err != nil {
		t.Fatalf("parseURL returned error: %v", err)
	}
	if domain != "example.com" {
		t.Errorf("Expected domain to be example.com, got '%s'", domain)
	}
	if scheme != "https" {
		t.Errorf("Expected scheme 'https', got '%s'", scheme)
	}
}

func TestResolveURL(t *testing.T) {
	baseURL := "http://example.com"
	domain := "example.com"
	scheme := "http"

	tests := []struct {
		href string
		expected string
		shouldPass bool
	}{
		{"/page1", "http://example.com/page1", true},
		{"http://example.com/page2", "http://example.com/page2", true},
		{"http://example.com/page3", "http://example.com/page3", true},
		{"http://external.com/page4", "", false},
		{"#anchor", "", false},
		{"", "", false},
	}

	for _, test := range tests {
		result, err := resolveURL(test.href, baseURL, domain, scheme)
		if test.shouldPass {
			if err != nil {
				t.Errorf("resoveURL(%s) returned error: %v", test.href, err)
			} else if result != test.expected {
				t.Errorf("resolveURL(%s) = %s; want %s", test.href, result, test.expected)
			}
		} else {
			if err == nil {
				t.Errorf("resolveURL(%s) should have failed", test.href)
			}
		}
	}
}


func TestExtractLinks(t *testing.T) {
	htmlPageContent := `
	<html>
		<body>
			<a href="/page1">Page 1</a>
			<a href="http://example.com/page2">Page 2</a>
			<a href="http://external.com/page3"> External Page </a>
			<a href="#anchorText">Anchor Text </a>
		</body>
	</html>
	`

	baseURL := "http://example.com"
	domain := "example.com"
	scheme := "http"

	doc, err := html.Parse(strings.NewReader(htmlPageContent))
	if err != nil {
		t.Fatalf("Error parsing HTML: %v", err)
	}

	links := extractLinks(doc, baseURL, domain, scheme)
	expectedLinks := []string{
		"http://example.com/page1",
		"http://example.com/page2",
	}

	if len(links) != len(expectedLinks) {
		t.Errorf("Expected %d links, got %d", len(expectedLinks), len(links))
	}

	linkSet := make(map[string]struct{})
	for _, link := range links {
		linkSet[link] = struct{}{}
	}

	for _, expectedLink := range expectedLinks {
		if _, exists := linkSet[expectedLink]; !exists {
			t.Errorf("Expected link %s not found", expectedLink)
		}
	}
}

func TestSafeSet_Add(t *testing.T) {
	set := &URLSet{urls: make(map[string]struct{})}

	url1 := "http://example.com"
	url2 := "http://example.com/about"

	if !set.Add(url1) {
		t.Errorf("Expected %s to be added to the set", url1)
	}
	if set.Add(url1) {
		t.Errorf("Expected %s to not be added again", url1)
	}
	if !set.Add(url2) {
		t.Errorf("Expected %s to be added to the set", url2)
	}
}

func TestSafeSet_Concurrency(t *testing.T) {
	set := &URLSet{urls: make(map[string]struct{})}
	urls := []string{
		"http://example.com/page1",
		"http://examle.com/page2",
		"http://example.com/page3",
		"http://example.com/page4",
	}

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			set.Add(u)
		}(url)
	}
	wg.Wait()

	if len(set.urls) != len(urls) {
		t.Errorf("Expected set length %d, got %d", len(urls), len(set.urls))
	}
}