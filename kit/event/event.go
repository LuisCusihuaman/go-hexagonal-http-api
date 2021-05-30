package event

import (
	"context"
	"github.com/google/uuid"
	"time"
)

//go:generate mockery --case=snake --outpkg=eventmocks --output=eventmocks --name=Bus

// Bus defines the expected behaviour from an event bus.
type Bus interface {
	// Publish is the method used to publish new events.
	Publish(context.Context, []Event) error
	// Subscribe is the method used to subscribe new event handlers.
	Subscribe(Type, Handler)
}

// Handler defines the expected behaviour from an event handler
type Handler interface {
	Handle(context.Context, Event) error
}

// Type represents a domain event type.
type Type string

// Event represents a domain command.
type Event interface {
	ID() string
	AggregateID() string
	OcurredOn() time.Time
	Type() Type
}

type BaseEvent struct {
	eventID     string
	aggregateID string
	occurredOn  time.Time
}

func NewBaseEvent(aggregateID string) BaseEvent {
	return BaseEvent{
		eventID:     uuid.New().String(),
		aggregateID: aggregateID,
		occurredOn:  time.Now(),
	}
}

func (b BaseEvent) ID() string {
	return b.eventID
}

func (b BaseEvent) OcurredOn() time.Time {
	return b.occurredOn
}

func (b BaseEvent) AggregateID() string {
	return b.aggregateID
}
