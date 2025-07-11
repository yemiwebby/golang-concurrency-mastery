# The Newspaper Printing Press Challenge (Producer-Consumer)

## Project Overview

This project demonstrates the **Producer-Consumer concurrency model** using **Golang**.  
It simulates a **newspaper company** where:

- **Journalists (Producers)** write articles and submit them to the printing queue.
- **Printers (Consumers)** retrieve articles from the queue and print them.
- **A Shared Queue (Buffer)** manages articles before they are printed.

This challenge explores **real-world concurrency issues**, such as:

- **Resource synchronization**
- **Race conditions**
- **Blocking vs. Non-blocking operations**
- **Deadlock prevention**

## Challenge

Before checking the solution, try implementing the Producer-Consumer model yourself!

Your Task:
Simulate a newspaper printing press where:

Journalists (Producers) write articles and submit them to a queue.
Printers (Consumers) retrieve articles from the queue and print them.
A Shared Queue (Buffer) manages the flow of articles.

Bonus:
Implement a priority queue to handle urgent news first.
Introduce graceful shutdown to safely stop the printing process.
Optimize the queue size for efficient processing.
Once you've attempted it, compare your approach with the provided solution!

---

## **Features**

**Concurrency using Channels** – Producers & Consumers work in parallel  
 **Priority Queue for Breaking News** – Urgent news gets printed first  
 **Graceful Shutdown using Context** – Ensures safe termination of processes  
 **Timeout for Idle Printers** – Prevents infinite waiting

---

## ** Running the Project**

Ensure you have Go installed (`go version` to check).  
Then, run:

```sh
go run main.go
```

## Running Tests

To verify functionality, run:

```sh
go test -race -v .
```

This:

- Runs all unit tests
- Detects race conditions
- Displays verbose output for debugging

## Related Concepts

- Goroutines & Channels – Efficient concurrency in Go
- Mutex vs Channels – Trade-offs in data synchronization
- Worker Pools – Managing concurrent tasks efficiently
- Go Context Package – Graceful cancellation of processes

## Future Enhancements

- Metrics Collection – Track articles processed per second
- Logging System – Capture events for debugging
- Configurable Queue Size – Allow dynamic buffer adjustments

## Author & Contributions

- Created by [Oluyemi Olususi](https://github.com/yemiwebby)
- Open to contributions & discussions! Feel free to fork & improve.

What’s Next?
Now that this project is completed, I’ll be working on the next concurrency challenge!

Stay tuned for: 📄 Article coming soon!

Let’s keep pushing the boundaries of Go concurrency!

```

```
