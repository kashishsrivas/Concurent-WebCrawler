package orchestrator

import (
	"context"
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
		Jobs:    make(chan crawler.Job, 100),
	}
}

func (wp *WorkerPool) Start(
	ctx context.Context,
	workerFunc func(context.Context, crawler.Job),
) {

	for i := 0; i < wp.Workers; i++ {

		wp.WG.Add(1)

		go func() {

			defer wp.WG.Done()

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
		}()
	}
}

func (wp *WorkerPool) Wait() {
	wp.WG.Wait()
}
