# 🔢 Concurrent Prime Number Generator – Overview & Requirements

## 🏗 Project Overview

This project implements the Sieve of Eratosthenes in a concurrent way using Go’s goroutines and channels. The generator starts from 2 and continuously filters out multiples of each discovered prime by creating a pipeline of filter goroutines. This design not only efficiently produces prime numbers but also demonstrates key concepts in concurrent programming, such as pipelines, dynamic goroutine creation, and channel-based communication.

## 🎯 Challenge

Before diving into the implementation, try solving the problem yourself!

Your Task:
Implement a concurrent prime number generator in Go using goroutines and channels. The generator should:

✅ Start from 2 and dynamically filter out non-prime numbers.
✅ Use channels to pass numbers through different filtering stages.
✅ Continuously generate prime numbers until a stop condition is met.

Bonus:

Implement a graceful shutdown mechanism.
Add configurable limits to control the number of primes generated.
Once you've given it a shot, check out the existing implementation to compare approaches! 🚀

---

## **🔥 Features**

✅ **Concurrent Pipeline:**

- Each prime discovered spawns a new goroutine to filter its multiples.

✅ **Dynamic Filter Stages:**

- The pipeline automatically expands as new primes are discovered.

✅ **Channel-Based Communication:**

- Uses channels to seamlessly pass numbers between stages.

✅ **Scalable & Efficient:**

- Designed to run indefinitely or up to a configurable limit.

✅ **Graceful Shutdown:**

- Provides mechanisms to terminate the generator cleanly.

✅ **Detailed Logging:**

- Logs each stage of number filtering and prime discovery for easy debugging.

---

## **🚀 Running the Project**

Ensure you have Go installed (`go version` to check).  
Then, run:

```sh
go run main.go
```

This command starts the prime number generator, which outputs primes to the console.

- 🧪 Running Tests
  To verify functionality and check for race conditions, run:

```sh
go test -race -v .
```

This command:

- Runs all unit tests
- Detects race conditions
- Provides verbose output for debugging

## 🔗 Related Concepts

- 📌 Goroutines & Channels:
  Efficiently manage concurrent processes in Go.

- 📌 Pipelines:
  Structure your program as a sequence of processing stages.

- 📌 Sieve of Eratosthenes:
  A classical algorithm for generating prime numbers.

- 📌 Dynamic Goroutine Creation:
  Spawning new goroutines for each new filter stage.

- 📌 Graceful Shutdown:
  Techniques for cleanly terminating concurrent processes.

## 📸 Future Enhancements

- 📌 Configurable Limits:
  Allow users to set a maximum number to search for primes.

- 📌 Performance Metrics:
  Track how many numbers are processed per second.

- 📌 Output Options:
  Save primes to a file (e.g., JSON or CSV) for further analysis.

- 📌 Distributed Processing:
  Explore scaling the solution across multiple machines or processes.

## 💡 Author & Contributions

🛠 Created by Oluyemi Olususi
📢 Open to contributions & discussions! Feel free to fork, improve, and share feedback.
