package kitchen

import (
	"time"
)

type Customer struct {
	Node

	ExitMessage string

	item Item
	maxWaitTime time.Duration
	eatTime time.Duration

	storage *Storage

	stopSignal chan bool
}

func NewCustomer(name string, storage *Storage, item Item, maxWaitTime time.Duration, eatTime time.Duration) *Customer {
	return &Customer{
		Node{NodeTypeCustomer, name},
		"",
		item,
		maxWaitTime,
		eatTime,
		storage,
		make(chan bool),
	}
}

func (c *Customer) Run() {
	select {
	case <-c.storage.GetMeal(c.item):
	case <-time.After(c.maxWaitTime):
		c.ExitMessage = "I hate this place!"
		return
	case <-c.stopSignal:
		return
	}

	select {
	case <-time.After(c.eatTime):
		c.ExitMessage = "I love this burger"
	case <-c.stopSignal:
	}
}

func (c *Customer) Stop() {
	c.stopSignal <- true
}