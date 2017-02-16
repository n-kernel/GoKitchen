package kitchen

import (
	"time"
	"github.com/Jeroenimoo/GoKitchen/util"
)

type Cook struct {
	Node

	storage     *Storage
	bakeTime    time.Duration
	BurgerCount int

	stopSignal chan bool
}

func NewCook(name string, storage *Storage, bakeTime time.Duration) *Cook {
	return &Cook{Node{NodeTypeCook, name}, storage, bakeTime, 0, make(chan bool, 1)}
}

func (c *Cook) Start() {
	for {
		select {
		case <-c.stopSignal:
		default:
			c.assembleBurger()
		}
	}
}

func (c *Cook) Stop() {
	c.stopSignal <- true
}

func (c *Cook) assembleBurger() {
	c.updateStatus(Waiting)
	// Wait for burger items to be available
	select {
	case <-util.Merge(c.storage.Get(Bread), c.storage.Get(Cheese), c.storage.Get(Tomato), c.storage.Get(Lettuce)):
	case <-c.stopSignal:
		return
	}

	c.updateStatus(Working)

	// Wait for burger preparation time
	time.Sleep(time.Second)

	c.updateStatus(Finished)
	c.storage.Burger <- true
}
