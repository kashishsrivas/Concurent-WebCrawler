package orchestrator

import (
	"context"
	"fmt"
	"sync"

	"Concurent-WebCrawler/internal/crawler"
)

type WorkerPool struct {
	Workers int
	Jobs    chan crawler.Job
	WG      sync.WaitGroup
}

func NewWorkerPool(workers int) *WorkerPool {

	return &WorkerPool{
		Workers: workers,
		Jobs:    make(chan crawler.Job, 5000),
	}
}

func (wp *WorkerPool) Start(
	ctx context.Context,
	workerFunc func(context.Context, crawler.Job),
) {

	for i := 0; i < wp.Workers; i++ {

		wp.WG.Add(1)

		go func(workerID int) {

			defer wp.WG.Done()

			fmt.Printf(
				"Worker %d started\n",
				workerID,
			)

			for {

				select {

				case <-ctx.Done():
					return

				case job, ok := <-wp.Jobs:

					if !ok {
						return
					}

					workerFunc(
						ctx,
						job,
					)
				}
			}
		}(i + 1)
	}
}

func (wp *WorkerPool) Wait() {
	wp.WG.Wait()
}
