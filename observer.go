package parallel

import (
	"log"
)

type Observer interface {
	PreStart(workerId string, job Job)
	PostStart(workerId string, job Job)
}

type TraceObserver struct {
}

func (ob *TraceObserver) PreStart(workerId string, job Job) {
	log.Println(workerId, "Starts", job.GetId())
}

func (ob *TraceObserver) PostStart(workerId string, job Job) {
	log.Println(workerId, "Finished", job.GetId())
}
