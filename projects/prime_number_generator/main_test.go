package main

import (
	"reflect"
	"testing"
)

func TestSieve(t *testing.T) {
	expected := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	result := Sieve(10)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Sieve(10) = %v; want %v", result, expected)
	}
}
