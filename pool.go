package parallel

import (
	"fmt"
	"sync"
)

type Job interface {
	GetId() string
	Start() error
}

type Pool interface {
	Start()
	Done()
	Wait()
	Add(job Job)
}

func CreatePool(size, maxWorker int, observers ...Observer) Pool {
	return &pool{
		jobChan:   make(chan Job, size),
		maxWorker: maxWorker,
		observers: observers,
	}
}

type pool struct {
	jobChan   chan Job
	maxWorker int
	wg        sync.WaitGroup
	observers []Observer
}

func (pool *pool) Add(job Job) {
	pool.jobChan <- job
}

func (pool *pool) Start() {
	for i := 0; i < pool.maxWorker; i++ {
		pool.wg.Add(1)
		go func(workerId string) {
			defer func() {
				pool.wg.Done()
			}()
			for {
				job, ok := <-pool.jobChan
				if !ok {
					return
				}
				for _, ob := range pool.observers {
					ob.PreStart(workerId, job)
				}
				job.Start()
				for _, ob := range pool.observers {
					ob.PostStart(workerId, job)
				}
			}
		}(fmt.Sprintf("Worker_%d", i))
	}
}

func (pool *pool) Done() {
	close(pool.jobChan)
}

func (pool *pool) Wait() {
	pool.wg.Wait()
}
