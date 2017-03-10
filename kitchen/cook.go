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

func (c *Cook) grabIngredients() <-chan bool {
	return util.Merge(c.storage.GetIngredient(Bread), c.storage.GetIngredient(Cheese), c.storage.GetIngredient(Tomato), c.storage.GetIngredient(Lettuce))
}

func (c *Cook) assembleBurger() {
	// Wait for burger items to be available
	grabChannel := c.grabIngredients()
	select {
		// Check if ingredients are available already
		case <-grabChannel:
		// If not wait for them and update status to waiting
		default:
			c.updateStatus(Waiting, "Missing ingredient(s)")

			// Wait until items are available or until the stop signal
			select {
			case <-grabChannel:
			case <-c.stopSignal:
				return
			}
	}

	c.updateStatus(Working, "Making burger")

	// Wait for burger preparation time
	time.Sleep(time.Second)

	select {
	// Add burger if there is any space
	case c.storage.GetMeal(Burger) <- true:
	// If not update status to waiting, and wait for space
	default:
		c.updateStatus(Waiting, "No customer!")
		c.storage.GetMeal(Burger) <- true
	}

	c.updateStatus(Finished, "Served burger")
}
