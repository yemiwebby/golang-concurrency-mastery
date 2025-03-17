package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	id int
	name string
	leftFork *Fork
	rightFork *Fork
}

func NewPhilosopher(id int, name string, leftFork, rightFork *Fork) *Philosopher {
	return &Philosopher{
		id: id,
		name: name,
		leftFork: leftFork,
		rightFork: rightFork,

	}
}

type Fork struct {
	id int
	mu sync.Mutex
}

func (p *Philosopher) dine(wg *sync.WaitGroup, arbitrator chan struct{}) {
	defer wg.Done()

	cycleOfEating := 3

	for i := 0; i < cycleOfEating; i++ {
		p.think()

		// Request permission before attempting to eat
		arbitrator <- struct{}{} // acquire permission 
		p.eat()
		<-arbitrator // release permission after eating
	}
}


func (p *Philosopher) think() {
	fmt.Printf("%s is thinking...\n", p.name)
	time.Sleep(time.Millisecond * time.Duration(100+10*p.id))
	fmt.Println("==================================")
}

func (p *Philosopher) eat() {
	// pick up the left fork
	p.leftFork.mu.Lock()
	fmt.Printf("%s picked up left fork %d\n", p.name, p.leftFork.id)

	// pick up the right fork
	p.rightFork.mu.Lock()
	fmt.Printf("%s picked up the right fork %d\n", p.name, p.rightFork.id)

	// simulate eating
	fmt.Printf("%s is eating...\n", p.name)
	time.Sleep(time.Millisecond * time.Duration(100+20*p.id))

	// put down the right fork
	p.rightFork.mu.Unlock()
	fmt.Printf("%s put down the right fork %d\n", p.name, p.rightFork.id)


	// put down the left fork
	p.leftFork.mu.Unlock()
	fmt.Printf("%s put down the left fork %d\n", p.name, p.leftFork.id)
}


func main() {
	numPhilosophers := 5

	forks := make([]*Fork, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		forks[i] = &Fork{id: i}
	}

	names := []string{"Socrates", "Plato", "Aristotle", "Confucius", "Nietzsche"}
	philosophers := make([]*Philosopher, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		philosophers[i] = &Philosopher{
			id: i,
			name: names[i],
			leftFork: forks[i],
			rightFork: forks[(i+1)%numPhilosophers],
		}
	}

	arbitrator := make(chan struct{}, numPhilosophers-1)

	var wg sync.WaitGroup
	for _, p := range philosophers {
		wg.Add(1)

		go p.dine(&wg, arbitrator)
	}

	wg.Wait()
	fmt.Println("Dinner is over!")
}