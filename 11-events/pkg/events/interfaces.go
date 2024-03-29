package events

import (
	"sync"
	"time"
)

type Event interface {
	GetName() string
	GetTimestamp() time.Time
	GetPayload() interface{}
}

type EventHandler interface {
	Handle(event Event, wg *sync.WaitGroup)
}

type EventDispatcher interface {
	Register(eventName string, handler EventHandler) error
	Unregister(eventName string, handler EventHandler) error
	Has(eventName string, handler EventHandler) bool
	Dispatch(event Event) error
	Clear()
}
