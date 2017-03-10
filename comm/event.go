package comm

type Event struct {
	Name string
	Data map[string]interface{}
}

type EventBus struct {
	Publish chan<- *Event
	publish <-chan *Event

	listeners map[chan<- *Event] bool

	Joining chan <-chan *Event
	Leaving chan <-chan *Event

	joining chan chan *Event
	leaving chan chan *Event
}

func NewEventBus() *EventBus {
	publish := make(chan *Event, 100)

	joining := make(chan chan *Event)
	leaving := make(chan chan *Event)


	return &EventBus{
		publish,
		publish,
		make(map[chan<- *Event] bool),
		joining,
		leaving,
		joining,
		leaving,
	}
}

func (nc *EventBus) Run() {
	for {
		select {
		case listener := <- nc.joining:
			nc.listeners[listener] = true
		case leaving := <- nc.leaving:
			delete(nc.listeners, leaving)
		case data := <- nc.publish:
			for listener := range nc.listeners {
				listener <- data
			}
		}
	}
}
