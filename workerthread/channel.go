package workerthread

import "fmt"

const (
	MaxQueue = 100
)

type Channel struct {
	ch         chan *Request
	thread     int
	threadPool []*Worker
	queue      chan *Request
}

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

func (ch *Channel) StartWorkers() {
	for _, worker := range ch.threadPool {
		worker.Start()
	}
}

func (ch *Channel) TakeRequest() *Request {
	return <-ch.queue
}

func (ch *Channel) PutRequest(req *Request) {
	ch.queue <- req
}
