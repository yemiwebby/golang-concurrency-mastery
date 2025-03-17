package main

import (
	"context"
	"math/rand"
	"sync"
	"time"

	"github.com/fatih/color"
)

/*
PRODUCER-CONSUMER
A conceptual challenge: "Newspaper Printing Press"

Scenario:
- **Producers (Journalists)** write articles and place them in the **printing queue**.
- **Consumers (Printers)** retrieve articles and print them.
- **Concurrency Challenges**: Prevent race conditions, prioritize breaking news, and handle delays.

THE SOLUTION
  ✅  Producers (Journalists)
	- Write an article every few seconds and place it in the printing queue.
	- If the queue is full, they must wait before submitting another article.
  ✅  Consumers (Printers)
    - Continuously take articles from the printing queue and print them.
	- If the queue is empty, they must wait for new articles
  ✅ Shared Resource (Queue)
    - Acts as a buffer between journalists and printers.
	- Must have a fixed size to prevent infinite growth.
	- Should support concurrent access safely.
  ✅ Concurrency challenges to Solve
    - Prevent race conditions when accessignt he queue.
	- Ensure journalists don't overwrite existing articles.
	- Ensure printers don't pick up the same article twice.
  ✅  Optional Enhancements
    - Introduce priorities (e.g, breaking news gets printed first)
	- Add delays (some articles take longer to print).
	- Simulate multiple printers & journalists

	Features:
  - Controlled **concurrency** using **channels**.
  - **Priority queue** for breaking news.
  - **Graceful shutdown** using `context.Context`.
  - **Timeout for printers** to avoid blocking.
*/

// Configurations
const (
	NumArticles = 10
	MaxQueueSize = 5
	NumPrinters = 2
	ArticleWriteTime = 2 
)

func newRandomGenerator() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func main() {
	// print out a welcome message
	color.Cyan("The Printing Press is OPEN!")
	color.Cyan("==============================================")

	// Shared queue for articles
	articleQueue := make(chan Article, MaxQueueSize)
	priorityQueue := make(chan Article, 2) 

	ctx, cancel := context.WithCancel(context.Background())

	// Initialize and Start Producer
	producer := NewProducer(ctx, articleQueue, priorityQueue, newRandomGenerator())
	go producer.Start()

	var wg sync.WaitGroup
	for i := 1; i <= NumPrinters; i++ {
		wg.Add(1)
		consumer := NewConsumer(i, articleQueue, priorityQueue, &wg, newRandomGenerator())
		go consumer.Start()
	}

	wg.Wait()

	color.Cyan("==============================================")
	color.Cyan("The Printing Press is CLOSED for the day!")
	cancel()
}