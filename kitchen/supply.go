package kitchen

import "time"

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
