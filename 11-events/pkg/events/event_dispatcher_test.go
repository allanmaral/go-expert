package events_test

import (
	"testing"
	"time"

	"github.com/allanmaral/go-expert/11-events/pkg/events"
)

type event struct {
	Payload interface{}
	Name    string
}

type eventHandler struct {
	ID int
}

var _ events.Event = (*event)(nil)
var _ events.EventHandler = (*eventHandler)(nil)

func (e *event) GetName() string {
	return e.Name
}

func (e *event) GetTimestamp() time.Time {
	return time.Now()
}

func (e *event) GetPayload() interface{} {
	return e.Payload
}

func (h *eventHandler) Handle(event events.Event) {
}

func TestRegister(t *testing.T) {
	t.Run("add handler when handler not already registered", func(t *testing.T) {
		dispatcher := events.NewEventDispatcher()
		handler := eventHandler{}
		eventName := "specificEventName"

		err := dispatcher.Register(eventName, &handler)

		assertNilError(t, err)
		assertDispatcherHasHandler(t, dispatcher, eventName, &handler)
	})
	t.Run("register multiple handlers when registering different handlers for the same event name", func(t *testing.T) {
		dispatcher := events.NewEventDispatcher()
		firstHandler := eventHandler{}
		secondHandler := eventHandler{}
		eventName := "sameEventName"

		err1 := dispatcher.Register(eventName, &firstHandler)
		err2 := dispatcher.Register(eventName, &secondHandler)

		assertNilError(t, err1)
		assertNilError(t, err2)
		assertDispatcherHasHandler(t, dispatcher, eventName, &firstHandler)
		assertDispatcherHasHandler(t, dispatcher, eventName, &firstHandler)
	})
	t.Run("returns error when registering handler twice", func(t *testing.T) {
		dispatcher := events.NewEventDispatcher()
		handler := eventHandler{}
		eventName := "anotherSpecificEventName"

		_ = dispatcher.Register(eventName, &handler)
		err := dispatcher.Register(eventName, &handler)

		assertCorrectError(t, err, events.ErrHandlerAlreadyRegistered)
	})
}

func assertDispatcherHasHandler(t testing.TB, dispatcher events.EventDispatcher, eventName string, handler events.EventHandler) {
	t.Helper()
	if !dispatcher.Has(eventName, handler) {
		t.Errorf("expected dispatcher to have registered handler")
	}
}

func assertNilError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("expected no error, got %q instead", got)
	}
}

func assertCorrectError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("expected error %q, got %q instead", want, got)
	}
}
