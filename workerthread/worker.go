package workerthread

// Worker is a worker thread object of worker thread pattern.
type Worker struct {
	Name    string
	channel Channel
}

// NewWorker returns worker object.
func NewWorker(name string, channel Channel) *Worker {
	return &Worker{name, channel}
}

// Start fetch request from queue and run.
func (w *Worker) Start() {
	go func() {
		for {
			req := w.channel.TakeRequest()
			req.Execute(w)
		}
	}()
}
