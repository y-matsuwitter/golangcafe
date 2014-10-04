package workerthread

import "fmt"

const (
	// MaxQueue is a queue size
	MaxQueue = 100
)

// Channel IS NOT A GO CHANNEL!!!!!
type Channel struct {
	ch         chan *Request
	thread     int
	threadPool []*Worker
	queue      chan *Request
}

// NewChannel creates channels
func NewChannel(threads int) *Channel {
	channel := Channel{
		make(chan *Request),
		threads,
		make([]*Worker, threads),
		make(chan *Request, MaxQueue),
	}
	for i := 0; i < threads; i++ {
		channel.threadPool[i] = NewWorker(fmt.Sprintf("Worker-%d", i), channel)
	}
	return &channel
}

// StartWorkers runs each thread
func (ch *Channel) StartWorkers() {
	for _, worker := range ch.threadPool {
		worker.Start()
	}
}

// TakeRequest return request.
func (ch *Channel) TakeRequest() *Request {
	return <-ch.queue
}

// PutRequest enqueue request.
func (ch *Channel) PutRequest(req *Request) {
	ch.queue <- req
}
