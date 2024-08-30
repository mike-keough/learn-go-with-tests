package sync

import "sync"

type Counter struct {
	mu        sync.Mutex
	Increment int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Lock()
	c.Increment++
}

func (c *Counter) Value() int {
	return c.Increment
}

func NewCounter() *Counter {
	return &Counter{}
}
