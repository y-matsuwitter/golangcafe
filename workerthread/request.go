package workerthread

import (
	"fmt"
	"math/rand"
	"time"
)

type Request struct {
	name   string
	number int
	rand   *rand.Rand
}

func NewRequest(name string, number int) *Request {
	return &Request{
		name,
		number,
		rand.New(rand.NewSource(1)),
	}
}

func (req *Request) Execute(worker *Worker) {
	fmt.Printf("%s executes %s\n", worker.Name, req)
	time.Sleep(time.Duration(req.rand.Float32()*1000.0) * time.Millisecond)
}

func (req *Request) String() string {
	return fmt.Sprintf("[ Request from %s No.%d ]", req.name, req.number)
}
