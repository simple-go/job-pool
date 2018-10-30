# parallel
Simple way to have parallel working workers

## Getting Started

### Usage
```
poolSize := 5
maxWorker := 2
pool := parallel.CreatePool(poolSize, maxWorker)
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
```

```
=== RUN   TestJobPool
2018/10/31 09:07:28 Worker_1 Starts Job_0
2018/10/31 09:07:28 Worker_0 Starts Job_1
2018/10/31 09:07:29 Worker_1 Finished Job_0
2018/10/31 09:07:29 Worker_1 Starts Job_2
2018/10/31 09:07:29 Worker_0 Finished Job_1
2018/10/31 09:07:29 Worker_0 Starts Job_3
2018/10/31 09:07:30 Worker_1 Finished Job_2
2018/10/31 09:07:30 Worker_1 Starts Job_4
2018/10/31 09:07:30 Worker_0 Finished Job_3
2018/10/31 09:07:30 Worker_0 Starts Job_5
2018/10/31 09:07:31 Worker_0 Finished Job_5
2018/10/31 09:07:31 Worker_1 Finished Job_4
2018/10/31 09:07:31 Worker_1 Starts Job_7
2018/10/31 09:07:31 Worker_0 Starts Job_6
2018/10/31 09:07:32 Worker_1 Finished Job_7
2018/10/31 09:07:32 Worker_1 Starts Job_8
2018/10/31 09:07:32 Worker_0 Finished Job_6
2018/10/31 09:07:32 Worker_0 Starts Job_9
2018/10/31 09:07:33 Worker_1 Finished Job_8
2018/10/31 09:07:33 Worker_0 Finished Job_9
--- PASS: TestJobPool (5.02s)
PASS

```