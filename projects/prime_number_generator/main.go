package main

import (
	"flag"
	"fmt"
)


func Sieve(n int) []int {
	ch := make(chan int)
	go Generate(ch) // Start generating integers starting from 2

	primes := make([]int, 0, n)
	for i := 0; i < n; i++ {
		prime := <-ch        
		primes = append(primes, prime)
		ch1 := make(chan int) 
		go Filter(ch, ch1, prime)
		ch = ch1 
	}
	return primes
}

func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}


func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func main() {
	n := flag.Int("n", 10, "number of primes to generate")
	flag.Parse()

	primes := Sieve(*n)
	for _, p := range primes {
		fmt.Println(p)
	}
}
