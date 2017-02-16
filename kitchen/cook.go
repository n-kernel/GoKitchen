package kitchen

import (
	"time"
	"fmt"
	"github.com/Jeroenimoo/GoKitchen/util"
)

type cook struct {
	storage     storage
	bakeTime    time.Duration
	BurgerCount int

	stopSignal chan bool
}

func NewCook(storage storage, bakeTime time.Duration) cook {
	return cook{storage,bakeTime, 0, make(chan bool, 1)}
}

func (c cook) Start() {
	for {
		select {
		case <-c.stopSignal:
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
