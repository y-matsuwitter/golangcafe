package workerthread

type Worker struct {
	Name    string
	channel Channel
}

func NewWorker(name string, channel Channel) *Worker {
	return &Worker{name, channel}
}

func (w *Worker) Start() {
	go func() {
		for {
			req := w.channel.TakeRequest()
			req.Execute(w)
		}
	}()
}
