package main

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestProducer_ProducesArticles(t *testing.T) {
	articleQueue := make(chan Article, 10)
	priorityQueue := make(chan Article, 2)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rnd := newRandomGenerator()

	var wg sync.WaitGroup
	wg.Add(1)

	producer := NewProducer(ctx, articleQueue, priorityQueue, rnd)

	go func() {
		defer wg.Done()
		producer.Start()
	}()

	wg.Wait()

	select {
	case article := <-articleQueue:
		assert.NotEmpty(t, article.Title, "Article should have a title")
	case <-time.After(1 * time.Second):
		t.Fatal("Producer did not generate any articles")
	}

	select {
	case article := <-priorityQueue:
		assert.True(t, article.Priority, "Priority article should be in priorityQueue")
	case <-time.After(1 * time.Second):
		t.Log("No priority articles were produced, which is fine.")
	}
}