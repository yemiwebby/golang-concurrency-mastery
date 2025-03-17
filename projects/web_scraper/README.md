# 🌐 Parallel Web Scraper – Overview & Requirements

## 🏗 Project Overview

This project is about building a concurrent web crawler in Go using worker pools to efficiently fetch and process web pages in parallel. The crawler starts from a given URL and explores all internal links on the site while:

- **Respecting rate limits:** Controlling the number of requests per second.
- **Deduplicating URLs:** Ensuring each page is visited only once using a thread-safe URL set.
- **Using buffered channels:** Decoupling URL production from consumption for smoother processing.
- **Handling errors gracefully:** Logging issues and avoiding deadlocks.

This challenge demonstrates key aspects of parallelism and concurrency in Go, including the use of goroutines, channels, mutexes, and worker pools.

## 🎯 Challenge

Before checking out the solution, try building your own parallel web scraper!

Your Task:
Create a concurrent web crawler that:

✅ Starts from a given URL and explores all internal links.
✅ Uses goroutines and channels to fetch pages in parallel.
✅ Implements rate limiting to control request frequency.
✅ Ensures each URL is visited only once using a thread-safe data structure.
✅ Gracefully handles errors and timeouts without crashing.

Bonus:
Implement depth control to limit how deep the crawler explores links.
Store results in a JSON or CSV file instead of just printing to the console.
Optimize performance by tuning the worker pool size.
Once you've built your own version, compare it with the existing implementation to refine your approach! 🚀

---

## **🔥 Features**

✅ **Concurrent Crawling:** Uses goroutines and channels to fetch and process pages in parallel  
✅ **Rate Limiting:** Implements a ticker to control the frequency of HTTP requests  
✅ **URL Deduplication:** Maintains a unique set of URLs to avoid redundant work  
✅ **Buffered URL Queue:** A channel buffer (e.g., capacity 1000) smooths out bursts in URL generation  
✅ **Robust Error Handling & Logging:** Detailed logs for debugging and error tracking  
✅ **Graceful Shutdown:** Ensures the crawler exits cleanly after processing all pages

---

## **🚀 Running the Project**

Ensure you have Go installed (`go version` to check).  
Then, run:

```sh
go run main.go <starting_url>
```

For example:

```sh
go run main.go https://example.com
```

## 🧪 Running Tests

To verify functionality, run:

```sh
go test -race -v .
```

This command:

- Runs all unit tests
- Detects race conditions
- Provides verbose output for debugging

## 🔗 Related Concepts

- 📌 Goroutines & Channels: Foundations of concurrent programming in Go
- 📌 Worker Pools: Managing a pool of workers for processing tasks efficiently
- 📌 Rate Limiting: Techniques to control the frequency of operations
- 📌 Mutexes & Sync Mechanisms: Ensuring safe concurrent access to shared data
- 📌 HTTP Clients & HTML Parsing: Fetching and processing web content using Go

## 📸 Future Enhancements

- 📌 Depth Control: Limit crawling to a specific depth or a maximum number of pages
- 📌 Output Storage: Save crawled data to JSON, CSV, or a database for later analysis
- 📌 Distributed Crawling: Scale the crawler across multiple machines or services
- 📌 Enhanced Error Handling: Implement robust retry and fallback strategies
- 📌 Customizable Configuration: Use command-line flags for adjustable parameters (e.g., rate limit, worker count)

## 💡 Author & Contributions

🛠 Created by Oluyemi Olususi
📢 Open to contributions & discussions! Feel free to fork, improve, and share feedback.

🚀 Let’s keep pushing the boundaries of Go concurrency!
