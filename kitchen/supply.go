package kitchen

import "time"

type Supply struct {
	storage Storage
	item    Item
	Delay   time.Duration

	stopSignal chan bool
}

func NewSupply(storage Storage, item Item, delay time.Duration) Supply {
	return Supply{storage, item, delay, make(chan bool)}
}

func (s Supply) Start() {
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

func (s Supply) Stop() {
	s.stopSignal <- true
}

func (s Supply) Item() Item {
	return s.item
}
