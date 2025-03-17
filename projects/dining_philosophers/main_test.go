package main

import (
	"testing"
	"time"
)

func TestDiningPhilosophers(t *testing.T) {
	done := make(chan struct{})

	go func() {
		main()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(10 * time.Second):
		t.Fatal("Dining philosophers simulation did not complete within 10 seconds (possible deadlock)")
	}
}
