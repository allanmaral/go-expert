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

func (h *eventHandler) Handle(event events.Event) {}

func TestEventDispatcher_Register(t *testing.T) {
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

func TestEventDispatcher_Clear(t *testing.T) {
	t.Run("remove handler when handler are present", func(t *testing.T) {
		dispatcher := events.NewEventDispatcher()
		handler := eventHandler{}
		eventName := "anyEventName"
		_ = dispatcher.Register(eventName, &handler)

		dispatcher.Clear()

		hasHandler := dispatcher.Has(eventName, &handler)
		if hasHandler {
			t.Error("expected event handler to be cleared")
		}
	})
}

func TestEventDispatcher_Has(t *testing.T) {
	t.Run("return false when there are no handlers registered", func(t *testing.T) {
		dispatcher := events.NewEventDispatcher()
		handler := eventHandler{}

		hasHandler := dispatcher.Has("anyEvent", &handler)

		if hasHandler {
			t.Error("expected dispatcher to have no handler")
		}
	})
	t.Run("return false when called with a different event name", func(t *testing.T) {
		dispatcher := events.NewEventDispatcher()
		handler := eventHandler{}
		dispatcher.Register("aSpecificEventName", &handler)

		hasHandler := dispatcher.Has("anotherEventName", &handler)

		if hasHandler {
			t.Error("expect dispatcher not to have handlers for \"anotherEventName\"")
		}
	})
	t.Run("return false when called with a different handler", func(t *testing.T) {
		dispatcher := events.NewEventDispatcher()
		registeredHandler := eventHandler{}
		unregisteredHandler := eventHandler{}
		dispatcher.Register("aSpecificEventName", &registeredHandler)

		hasHandler := dispatcher.Has("anotherEventName", &unregisteredHandler)

		if hasHandler {
			t.Error("expected only one handler to be registered")
		}
	})
	t.Run("return true when there is a handler registered with the same event name", func(t *testing.T) {
		dispatcher := events.NewEventDispatcher()
		handler := eventHandler{}
		eventName := "anotherSpecificEventName"
		dispatcher.Register(eventName, &handler)

		hasHandler := dispatcher.Has(eventName, &handler)

		if !hasHandler {
			t.Error("expected handler to be registered")
		}
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
