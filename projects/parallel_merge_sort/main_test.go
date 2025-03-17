package main

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestParallelMergeSort(t *testing.T) {
	// Test with a fixed input.
	input := []int{5, 3, 8, 4, 2, 7, 1, 6, 0, 9}
	result := parallelMergeSort(input)
	expected := make([]int, len(input))
	copy(expected, input)
	sort.Ints(expected)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("parallelMergeSort(%v) = %v; want %v", input, result, expected)
	}
}

func BenchmarkParallelMergeSort(b *testing.B) {
	size := 10000
	arr := make([]int, size)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range arr {
		arr[i] = r.Intn(100000)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = parallelMergeSort(arr)
	}
}

func TestSequentialVsParallel(t *testing.T) {
	size := 10000
	arr := make([]int, size)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range arr {
		arr[i] = r.Intn(100000)
	}

	seq := sequentialMergeSort(arr)
	par := parallelMergeSort(arr)

	if !reflect.DeepEqual(seq, par) {
		t.Error("Sequential and parallel merge sort results differ")
	}
}