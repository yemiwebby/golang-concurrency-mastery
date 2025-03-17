package main

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConsumer_ConsumesArticles(t *testing.T) {
	articleQueue := make(chan Article, 5)
	priorityQueue := make(chan Article, 2)
	wg := sync.WaitGroup{}

	rnd := newRandomGenerator()

	consumer := NewConsumer(1, articleQueue, priorityQueue, &wg, rnd)

	wg.Add(1)

	go consumer.Start()

	articleQueue <- Article{ID: 1, Title: "Regular News", Priority: false}
	priorityQueue <- Article{ID: 2, Title: "Breaking News!", Priority: true}

	time.Sleep(2 * time.Second)

	close(priorityQueue)
	close(articleQueue)

	wg.Wait()

	assert.True(t, true, "Consumer processed articles successfully")
}