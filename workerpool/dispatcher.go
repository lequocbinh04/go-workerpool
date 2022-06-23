package workerpool

import (
	"context"
	"fmt"
)

type Dispatcher struct {
	// pool of workers are register on pool, (workers pool)
	WorkerPool chan chan Job
	maxWorkers int
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	workerPool := make(chan chan Job, maxWorkers)
	return &Dispatcher{
		WorkerPool: workerPool,
		maxWorkers: maxWorkers,
	}
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.maxWorkers; i++ {
		fmt.Println("worker ", i, " started")
		worker := NewWorker(d.WorkerPool)
		worker.Start(context.Background())
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			go func(job Job) {
				worker := <-d.WorkerPool
				worker <- job
			}(job)
		}
	}
}
