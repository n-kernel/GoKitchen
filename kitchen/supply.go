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
		// Wait for work delay to complete or stop if signaled to
		select {
		case <-time.After(s.Delay):
		case <-s.stopSignal:
			return
		}

		select {
		// Add ingredient if not full
		case s.storage.GetIngredient(s.item) <- true:
		// Else update status and wait for space or stop signal
		default:
			s.updateStatusMessage(Waiting, "Storage full")

			// Wait for space to add ingredient or until the stop signal
			select {
			case s.storage.GetIngredient(s.item) <- true:
			case <-s.stopSignal:
				return
			}
		}
		s.updateStatus(Finished)
	}
}

func (s Supply) Stop() {
	s.stopSignal <- true
}

func (s Supply) Item() Item {
	return s.item
}
