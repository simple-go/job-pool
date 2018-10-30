package parallel_test

import (
	"fmt"
	"parallel"
	"testing"
	"time"
)

func TestJobPool(t *testing.T) {
	poolSize := 5
	maxWorker := 2
	pool := parallel.CreatePool(poolSize, maxWorker, &parallel.TraceObserver{})
	pool.Start()
	go func() {
		for i := 0; i < 10; i++ {
			jobId := fmt.Sprintf("Job_%d", i)
			job := parallel.NewClosureJob(jobId, func() error {
				time.Sleep(1 * time.Second)
				return nil
			})
			pool.Add(job)
		}
		pool.Done()
	}()
	pool.Wait()
}
