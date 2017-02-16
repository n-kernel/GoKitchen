package kitchen

import (
	"time"
	"fmt"
	"github.com/Jeroenimoo/GoKitchen/util"
)

type item int
const (
	Bread item = iota
	Cheese
	Tomato
	Lettuce

	Burger
)

var Items = []item{
	Bread,
	Cheese,
	Tomato,
	Lettuce,

	Burger,
}

// Holds buffers of ingredients
type storage struct {
	Ingredients map[item] chan bool

	Burger chan bool
}

func NewStorage() storage {
	s := storage{}
	s.Ingredients = make(map[item] chan bool)

	for _, item := range Items {
		s.Ingredients[item] = make(chan bool, 100)
	}

	s.Burger = make(chan bool, 100)

	return s
}

func (s storage) Get(item item) chan bool {
	return s.Ingredients[item]
}

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

type supply struct {
	storage storage
	item    item
	Delay   time.Duration

	stopSignal chan bool
}

func NewSupply(storage storage, item item, delay time.Duration) supply {
	return supply{storage, item, delay, make(chan bool)}
}

func (s supply) Start() {
	for {
		time.Sleep(s.Delay)
		select {
		case <-s.stopSignal:
			return
		default:
			s.storage.Get(s.item) <- true
		}
	}
}

func (s supply) Stop() {
	s.stopSignal <- true
}