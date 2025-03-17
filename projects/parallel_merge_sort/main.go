package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

// threshold defines the maximum array size to process sequentially.
const threshold = 100

func parallelMergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	if len(arr) <= threshold {
		return sequentialMergeSort(arr)
	}

	mid := len(arr) / 2
	var left, right []int
	var wg sync.WaitGroup
	wg.Add(2)

	// Sort left half concurrently.
	go func() {
		defer wg.Done()
		left = parallelMergeSort(arr[:mid])
	}()

	// Sort right half concurrently.
	go func() {
		defer wg.Done()
		right = parallelMergeSort(arr[mid:])
	}()

	wg.Wait()
	return merge(left, right)
}

func sequentialMergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := sequentialMergeSort(arr[:mid])
	right := sequentialMergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func main() {
	// Allow the user to specify the number of primes to generate via a flag.
	n := flag.Int("n", 10, "number of elements in the array")
	flag.Parse()

	// Generate a random array of integers.
	size := *n
	arr := make([]int, size)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range arr {
		arr[i] = r.Intn(10000)
	}

	start := time.Now()
	sorted := parallelMergeSort(arr)
	elapsed := time.Since(start)

	if len(sorted) > 20 {
		fmt.Println("Sorted array (first 20 elements):", sorted[:20])
	} else {
		fmt.Println("Sorted array:", sorted)
	}
	fmt.Printf("Sorting took %s\n", elapsed)

	expected := make([]int, len(arr))
	copy(expected, arr)
	sort.Ints(expected)
	if !sort.SliceIsSorted(sorted, func(i, j int) bool { return sorted[i] < sorted[j] }) {
		fmt.Println("The array is not sorted correctly.")
	}
}