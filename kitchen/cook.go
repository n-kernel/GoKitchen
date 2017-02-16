package kitchen

import (
	"time"
	"fmt"
	"github.com/Jeroenimoo/GoKitchen/util"
)

type Cook struct {
	storage     Storage
	bakeTime    time.Duration
	BurgerCount int

	stopSignal chan bool
}

func NewCook(storage Storage, bakeTime time.Duration) Cook {
	return Cook{storage, bakeTime, 0, make(chan bool, 1)}
}

func (c Cook) Start() {
	for {
		select {
		case <-c.stopSignal:
		default:
			c.assembleBurger()
		}
	}
}

func (c Cook) Stop() {
	c.stopSignal <- true
}

func (c Cook) assembleBurger() {
	fmt.Println("Started making a burger!")

	// Wait for burger items to be available
	select {
	case <-util.Merge(c.storage.Get(Bread), c.storage.Get(Cheese), c.storage.Get(Tomato), c.storage.Get(Lettuce)):
	case <-c.stopSignal:
		return
	}

	// Wait for burger preparation time
	time.Sleep(time.Second)

	c.storage.Burger <- true
	fmt.Println("Created a burger!")
}
