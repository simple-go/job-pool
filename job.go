package parallel

type closureJob struct {
	id          string
	closureFunc func() error
}

func (w *closureJob) GetId() string {
	return w.id
}

func (w *closureJob) Start() error {
	return w.closureFunc()
}

func NewClosureJob(id string, closureFunc func() error) Job {
	return &closureJob{
		id:          id,
		closureFunc: closureFunc,
	}
}
