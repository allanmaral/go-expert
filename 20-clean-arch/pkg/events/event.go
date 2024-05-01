package events

import (
	"sync"
	"time"
)

type Event interface {
	GetName() string
	GetTimestamp() time.Time
	GetPayload() interface{}
	SetPayload(payload interface{})
}

type EventHandler interface {
	Handle(event Event, wg *sync.WaitGroup)
}
