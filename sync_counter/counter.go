package main

import "sync"

type Counter struct {
	value int
	mu    sync.Mutex
}

func (c *Counter) Inc() {
	defer c.mu.Unlock()
	c.mu.Lock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
