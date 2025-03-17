package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

// Represents a news article
type Article struct {
	ID int
	Title string
	Author string
	Description string
	Priority bool
}

// Producer (Journalists) werites articles and sends to queue
type Producer struct {
	ctx context.Context
	articleQueue chan Article
	priorityQueue chan Article
	rnd *rand.Rand
}

func NewProducer(ctx context.Context, queue chan Article, priorityQueue chan Article, rnd *rand.Rand) *Producer {
	return &Producer{
		ctx: ctx,
		articleQueue: queue,
		priorityQueue: priorityQueue,
		rnd: rnd,
	}
}

// Start the producer
func (p *Producer) Start() {
	for i := 1; i <= NumArticles; i++ {

		select {
		case <-p.ctx.Done():
			color.Red("Producer received shutdown signal. Stopping article creation!")
			close(p.articleQueue)
			close(p.priorityQueue)
			return
		default:
			time.Sleep(time.Duration(p.rnd.Intn(ArticleWriteTime)+1) * time.Second)

			article := Article{
				ID: i,
				Title: fmt.Sprintf("Breaking News #%d", i),
				Author: "Journalist " + fmt.Sprint(p.rnd.Intn(5)+1),
			    Description: fmt.Sprintf("Article %d is ready for printing!", i),
				Priority: p.rnd.Intn(10) < 3, // 30% chance its breaking news
			}

			if article.Priority {
				select {
				case p.priorityQueue <- article:
					color.Magenta("[%s] Breaking News! %s", time.Now().Format("15:04:05"), article.Title)
				default:
					color.Red("Priority queue is full. Sending to regular queue.")
					p.articleQueue <- article
				}
			} else {
				p.articleQueue <- article
				color.Green("[%s] Journalist wrote: %s", time.Now().Format("15:04:05"), article.Title)
			}
		}
	}

	close(p.articleQueue)
	close(p.priorityQueue)
}