package workerthread

import (
	"math/rand"
	"time"
)

// Client is a client thread in worker thread pattern.
type Client struct {
	name    string
	channel *Channel
	rand    *rand.Rand
}

// NewClient creates client.
func NewClient(name string, channel *Channel) *Client {
	return &Client{
		name,
		channel,
		rand.New(rand.NewSource(1)),
	}
}

// Run is a runner for each request.
func (c *Client) Run() {
	go func() {
		i := 0
		for {
			req := NewRequest(c.name, i)
			c.channel.PutRequest(req)
			time.Sleep(time.Duration(c.rand.Float32()*1000.0) * time.Millisecond)
			i++
		}
	}()
}
