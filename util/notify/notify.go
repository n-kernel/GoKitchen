package notify

import (
	"encoding/json"
	"fmt"
)

type Notifier struct {
	notify chan []byte
	listeners map[chan<- []byte] bool

	Joining chan <-chan []byte
	Leaving chan <-chan []byte

	joining chan chan []byte
	leaving chan chan []byte
}

func NewNotifier() *Notifier {
	joining := make(chan chan [] byte)
	leaving := make(chan chan [] byte)

	return &Notifier{
		make(chan []byte),
		make(map[chan<- [] byte] bool),
		joining,
		leaving,
		joining,
		leaving,
	}
}

func (nc *Notifier) Run() {
	for {
		select {
		case listener := <- nc.joining:
			nc.listeners[listener] = true
		case leaving := <- nc.leaving:
			delete(nc.listeners, leaving)
		case data := <- nc.notify:
			for listener := range nc.listeners {
				listener <- data
			}
		}
	}
}

func (nc *Notifier) Message(data map[string]interface{}) {
	bData, err := json.Marshal(data)

	if err != nil {
		fmt.Errorf("err: %d", err)
		return
	}

	nc.notify <- bData
}
