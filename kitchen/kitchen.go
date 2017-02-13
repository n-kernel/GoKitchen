package kitchen

import (
	"time"
	"fmt"
)

// Holds buffers of ingredients
type Storage struct {
	Bread chan bool
	Cheese chan bool
	Tomato chan bool
	Lettuce chan bool

	Burger chan bool
}

type cook struct {
	storage Storage
	bakeTime time.Duration
	BurgerCount int
	stopSignal chan bool
}

func NewCook(storage Storage, bakeTime time.Duration) cook {
	return cook{storage,bakeTime, 0, make(chan bool, 1)}
}

func (c cook) Start() {
	for {
		select {
		case <-c.stopSignal:
			return
		default:
			c.assembleBurger()
		}
	}
}

func (c cook) Stop() {
	c.stopSignal <- true
}

func (c cook) assembleBurger() {
	fmt.Println("Started making a burger!")

	// Wait for burger items to be available
	<-c.storage.Bread
	<-c.storage.Cheese
	<-c.storage.Tomato
	<-c.storage.Lettuce

	// Wait for burger preparation time
	time.Sleep(time.Second)

	c.storage.Burger <- true
	fmt.Println("Created a burger!")
}