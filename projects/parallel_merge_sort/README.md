# 🔢 Parallel Merge Sort – Overview & Requirements

## 🏗 Project Overview

This project implements a merge sort algorithm in a concurrent manner using Go’s goroutines and sync primitives. The idea is to sort a slice of integers by dividing the array into subarrays, sorting them concurrently, and then merging the sorted subarrays. This solution demonstrates the power of parallel processing and the use of channels and wait groups for synchronization in Go.

## 🎯 Challenge

Before reviewing the solution, try implementing a parallel merge sort yourself!

Your Task:
Build a concurrent merge sort algorithm that:

✅ Uses goroutines to sort subarrays in parallel.
✅ Implements a sequential fallback for small arrays to reduce overhead.
✅ Uses sync.WaitGroup to manage goroutine execution.
✅ Efficiently merges sorted subarrays back into a final sorted list.

Bonus:
Tune the threshold value to balance concurrency and performance.
Add benchmarking tests to compare sequential vs. parallel performance.
Extend the implementation to sort generic types using Go generics.
Once you've attempted it, compare your approach with the provided solution and optimize further! 🚀

---

## **🔥 Features**

✅ **Parallel Processing:**

- Large subarrays are sorted concurrently using goroutines.

✅ **Sequential Fallback:**

- For small slices (below a defined threshold), a sequential merge sort is used to avoid unnecessary overhead.

✅ **Efficient Merging:**

- Merges sorted subarrays using an efficient merging function.

✅ **Concurrency Primitives:**

- Uses `sync.WaitGroup` to coordinate goroutines.
- Uses a threshold constant to balance overhead vs. parallelism.

✅ **Extensible & Easy to Understand:**

- Clean separation of functions for generating, sorting, and merging.
- Can be extended or modified to test performance on various data sizes.

---

## **🚀 Running the Project**

Ensure you have Go installed (`go version` to check).  
Then, run:

```sh
go run main.go
```

This will sort a randomly generated array (or you can modify main() to sort any input slice) and print the sorted result.

## 🧪 Running Tests

To verify functionality and check for race conditions, run:

```sh
go test -race -v .
```

This command:

- Runs all unit tests in the project.
- Detects race conditions.
- Provides verbose output for debugging.

## 🔗 Related Concepts

- 📌 Goroutines & Concurrency:
  Running functions concurrently to improve performance.

- 📌 Divide and Conquer:
  Splitting a problem into smaller subproblems, solving them independently, and merging results.

- 📌 Wait Groups:
  Synchronizing multiple goroutines.

- 📌 Merge Sort Algorithm:
  A classic sorting algorithm using a divide-and-conquer strategy.

## 📸 Future Enhancements

- 📌 Configurable Threshold:
  Allow users to set the cutoff value for switching to sequential sorting.

- 📌 Benchmarking:
  Add benchmarks to compare the performance of the parallel versus sequential implementation.

- 📌 Sorting Different Data Types:
  Generalize the algorithm to sort other types (using interfaces or generics in Go 1.18+).

## 💡 Author & Contributions

🛠 Created by Oluyemi Olususi
📢 Open to contributions & discussions! Feel free to fork, improve, and share feedback.
