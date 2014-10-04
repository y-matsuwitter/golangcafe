package workerthread

import (
	"math/rand"
	"time"
)

type Client struct {
	name    string
	channel *Channel
	rand    *rand.Rand
}

func NewClient(name string, channel *Channel) *Client {
	return &Client{
		name,
		channel,
		rand.New(rand.NewSource(1)),
	}
}

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
