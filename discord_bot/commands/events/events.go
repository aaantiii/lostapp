package events

import "time"

type Event struct {
	Kind      string
	Timestamp time.Time
	Data      interface{}
}

type EventBus struct {
	subscriber map[string][]chan<- Event
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscriber: make(map[string][]chan<- Event),
	}
}

func NewEvent(kind string, data any) Event {
	return Event{
		Kind:      kind,
		Timestamp: time.Now(),
		Data:      data,
	}
}

func (eb *EventBus) Subscribe(kind string, ch chan<- Event) {
	eb.subscriber[kind] = append(eb.subscriber[kind], ch)
}

func (eb *EventBus) Publish(event Event) {
	for _, subscriber := range eb.subscriber[event.Kind] {
		subscriber <- event
	}
}
