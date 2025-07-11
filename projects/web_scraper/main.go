package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html"
)

type URLSet struct {
	mu   sync.RWMutex
	urls map[string]struct{}
}

// Returns true if the URL was not seen before.
func (s *URLSet) Add(url string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.urls[url]; exists {
		return false
	}
	s.urls[url] = struct{}{}
	return true
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(len(os.Args))
		log.Println("[USAGE]: go run main.go <starting_url>")
		return
	}

	startingURL := os.Args[1]
	domain, scheme, err := parseURL(startingURL)
	if err != nil {
		fmt.Println("Invalid URL:", err)
		return
	}

	rPerSecond := 5

	startCrawling(startingURL, domain, scheme, rPerSecond)
}

func parseURL(rURL string) (string, string, error) {
	u, err := url.ParseRequestURI(rURL)
	if err != nil {
		return "", "", err
	}

	fmt.Println(u, "Starting URL")

	return u.Hostname(), u.Scheme, nil
}

func startCrawling(sURL, domain, scheme string, requestPerSecond int) {
	visited := &URLSet{urls: make(map[string]struct{})}
	var wg sync.WaitGroup
	urlQueue := make(chan string, 1000)
	done := make(chan struct{})
	rateLimiter := time.Tick(time.Second / time.Duration(requestPerSecond))

	numWorkers := 5
	for i := 0; i < numWorkers; i++ {
		go worker(urlQueue, visited, &wg, domain, scheme, rateLimiter, done)
	}

	// Add Starting URL.
	if visited.Add(sURL) {
		wg.Add(1)
		urlQueue <- sURL
		log.Println("[DEBUG] Starting URL added to queue:", sURL)
	}

	wg.Wait()
	close(urlQueue)
	close(done)
	log.Println("Done crawling. Total unique URLs visited:", len(visited.urls))

}

func worker(urlQueue chan string, visited *URLSet, wg *sync.WaitGroup, domain, scheme string, rateLimiter <-chan time.Time, done chan struct{}) {

	for {
		select {
		case <-done:
			log.Println("[INFO] Worker exiting...")
			return
		case currentURL, ok := <-urlQueue:
			if !ok {
				return
			}
			<-rateLimiter
			processURL(currentURL, urlQueue, visited, wg, domain, scheme)
		}
	}
}

func processURL(currentURL string, urlQueue chan string, visited *URLSet, wg *sync.WaitGroup, domain, scheme string) {
	defer wg.Done()

	resp, err := fetchURL(currentURL)
	if err != nil {
		log.Printf("[ERROR]: Unable to fetch URL  %s: %s", currentURL, err)
		return
	}

	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Printf("[ERROR]: Error parsing HTML from %s: %v\n", currentURL, err)
		return
	}

	links := extractLinks(doc, currentURL, domain, scheme)
	log.Printf("[INFO] Found %d links on %s", len(links), currentURL)

	for _, link := range links {
		log.Println("[DEBUG]: Extracted link:", link)
		if visited.Add(link) {
			select {
			case urlQueue <- link:
				wg.Add(1)
			default:
				log.Println("[WARNING]: Queue full, dropping:", link)
			}
		} else {
			log.Println("[DEBUG]: Duplicate link, skipped:", link)
		}
	}
}

func fetchURL(currentURL string) (*http.Response, error) {
	resp, err := http.Get(currentURL)
	if err != nil {
		log.Printf("[ERROR]: Failed to fetch %s: %v", currentURL, err)
		return nil, fmt.Errorf("error fetching %s: %v", currentURL, err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Printf("[ERROR]: Non-OK HTTP status %d for %s", resp.StatusCode, currentURL)
		resp.Body.Close()
		return nil, fmt.Errorf("Non-OK HTTP status for %s: %s", currentURL, resp.Status)
	}
	return resp, nil
}

func resolveURL(href, baseURL, domain, scheme string) (string, error) {
	href = strings.TrimSpace(href)
	if href == "" || strings.HasPrefix(href, "#") {
		log.Println("[WARNING]: Skipping empty or anchor link:", href)
		return "", fmt.Errorf("invalid or empty URL")
	}

	link, err := url.Parse(href)
	if err != nil {
		log.Println("[WARNING]: Could not parse URL:", href)
		return "", err
	}

	// Check if link is absolute or relative
	base, err := url.Parse(baseURL)
	if err != nil {
		log.Println("[ERROR]: Could not parse base URL:", baseURL)
		return "", err
	}

	resolvedLink := base.ResolveReference(link)

	// Ensure we only crawl within the same domain
	if resolvedLink.Hostname() != domain {
		log.Println("[WARNING]: Skipping external domain:", resolvedLink.String())
		return "", fmt.Errorf("link %s does not belong to the domain %s", resolvedLink.String(), domain)
	}

	if resolvedLink.Scheme == "" {
		resolvedLink.Scheme = scheme
	}

	log.Println("[INFO]: Resolved link:", resolvedLink.String())

	return resolvedLink.String(), nil
}

func extractLinks(node *html.Node, baseURL, domain, scheme string) []string {
	var links []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					link, err := resolveURL(attr.Val, baseURL, domain, scheme)
					if err == nil {
						links = append(links, link)
					} else {
						log.Println("[WARNING]: Ignored link:", attr.Val, "Error:", err)
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(node)
	log.Println("[INFO]: Extracted", len(links), "valid links from", baseURL)
	return links
}
