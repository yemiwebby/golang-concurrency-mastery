package main

import (
	"math/rand"
	"sync"
	"time"

	"github.com/fatih/color"
)

// Consumer (Printers) retrieves and prints articles
type Consumer struct {
	ID int
	articleQueue chan Article
	priorityQueue chan Article
	wg *sync.WaitGroup
	rnd *rand.Rand
}

// NewConsumer initializes a consumer
func NewConsumer(id int, queue chan Article, priorityQueue chan Article, wg *sync.WaitGroup, rnd *rand.Rand) *Consumer {
	return &Consumer{
		ID: id,
		articleQueue: queue,
		priorityQueue: priorityQueue,
		wg: wg,
		rnd: rnd,
	}
}

// start the consumer
func (c *Consumer) Start() {
	defer c.wg.Done()
	for {
		select {
		case article, ok := <-c.priorityQueue:
			if !ok {
				color.Yellow("Printer #%d detected priority queue closed. Exiting...", c.ID)
				return
			}
			color.Magenta("Printer #%d printing BREAKING NEWS: %s", c.ID, article.Title)

		case article, ok := <-c.articleQueue:
			if !ok {
				color.Yellow("Printer #%d detected regular queue closed. Exiting...", c.ID)
				return
			}
			color.Cyan("Printer #%d printing: %s", c.ID, article.Title)

		case <-time.After(5 * time.Second):
			color.Red("Printer #%d is idle... waiting for articles", c.ID)
		}

		time.Sleep(time.Duration(c.rnd.Intn(3)+1) * time.Second)
	}
}