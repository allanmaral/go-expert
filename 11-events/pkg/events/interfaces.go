package events

import (
	"time"
)

type Event interface {
	GetName() string
	GetTimestamp() time.Time
	GetPayload() interface{}
}

type EventHandler interface {
	Handle(event Event)
}

type EventDispatcher interface {
	Register(eventName string, handler EventHandler) error
	Unregister(eventName string, handler EventHandler) error
	Has(eventName string, handler EventHandler) bool
	Dispatch(event Event) error
	Clear() error
}
