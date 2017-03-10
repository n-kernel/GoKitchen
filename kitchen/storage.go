package kitchen

import (
	"time"
	"strconv"
)

const bufferSize = 10

// Holds buffers of ingredients
type Storage struct {
	Node

	Ingredients map[Item] chan bool

	Meals map[Item] chan bool
}

func NewStorage() *Storage {
	s := Storage{}
	s.Node = Node{NodeTypeStorage, "Storage"}

	s.Ingredients = make(map[Item] chan bool)

	for _, item := range Ingredients {
		s.Ingredients[item] = make(chan bool, bufferSize)
	}

	s.Meals = make(map[Item] chan bool)
	s.Meals[Burger] = make(chan bool)

	go func() {
		for {
			full := false
			message := ""

			for item := range s.Ingredients {
				bufLen := len(s.Ingredients[item])
				message += Items[item].GetName() + ": " + strconv.Itoa(bufLen) + " "

				if !full {
					full = bufLen == bufferSize
				}
			}

			var status Status
			if full {
				status = Waiting
			} else {
				status = Working
			}

			s.updateStatus(status, message)
			time.Sleep(time.Second)
		}
	}()

	return &s
}

func (s Storage) GetMeal(item Item) chan bool {
	return s.Meals[item]
}

func (s Storage) GetIngredient(item Item) chan bool {
	return s.Ingredients[item]
}
