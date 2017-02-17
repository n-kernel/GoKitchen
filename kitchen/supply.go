package kitchen

import (
	"time"
)

type Supply struct {
	Node

	storage *Storage
	item    Item
	Delay   time.Duration

	stopSignal chan bool
}

func NewSupply(name string, storage *Storage, item Item, delay time.Duration) *Supply {
	return &Supply{Node{NodeTypeSupply, name}, storage, item, delay, make(chan bool), }
}

func (s *Supply) Start() {
	for {
		s.updateStatus(Working)
		time.Sleep(s.Delay)
		select {
		case <-s.stopSignal:
			return
		default:
			s.updateStatus(Finished)
			s.storage.GetIngredient(s.item) <- true
		}
	}
}

func (s Supply) Stop() {
	s.stopSignal <- true
}

func (s Supply) Item() Item {
	return s.item
}