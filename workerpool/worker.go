package workerpool

import (
	"context"
	log "github.com/sirupsen/logrus"
)

type JobHandler func(ctx context.Context, args []interface{}) error

type Job struct {
	handler JobHandler
	Args    []interface{}
}

func NewJob(handler JobHandler) Job {
	return Job{
		handler: handler,
	}
}

var JobQueue chan Job

type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

func InitJobQueue() {
	JobQueue = make(chan Job)
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
	}
}

func (w Worker) Start(ctx context.Context) {
	go func() {
		for {
			// register the current worker into the worker queue.
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				if err := job.handler(ctx, job.Args); err != nil {
					log.Errorf("Error when doing job: %s", err.Error())
				}

			case <-w.quit:
				return
			}
		}
	}()
}

// Stop signals the worker to stop listening for work requests.
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
