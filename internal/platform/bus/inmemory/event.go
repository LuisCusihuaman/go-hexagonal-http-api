package inmemory

import (
	"context"
	"github.com/LuisCusihuaman/go-hexagonal-http-api/kit/event"
)

// EventBus is an in-memory implementation of the event.Bus.
type EventBus struct {
	// Map<EventType, []Handler> handlers
	handlers map[event.Type][]event.Handler
}

// NewEventBus initializes a new EventBus.
func NewEventBus() *EventBus {
	return &EventBus{
		handlers: make(map[event.Type][]event.Handler),
	}
}

// Publish implements the event.Bus interface.
func (b *EventBus) Publish(ctx context.Context, events []event.Event) error {
	for _, evt := range events {
		handlers, ok := b.handlers[evt.Type()]

		if !ok {
			return nil
		}

		for _, handler := range handlers {
			_ = handler.Handle(ctx, evt)
		}
	}
	return nil
}

// Subscribe implements the event.Bus interface.
func (b *EventBus) Subscribe(evtType event.Type, handler event.Handler) {
	suscribersForType, ok := b.handlers[evtType]
	if !ok {
		b.handlers[evtType] = []event.Handler{handler}
	}

	suscribersForType = append(suscribersForType, handler)
}
