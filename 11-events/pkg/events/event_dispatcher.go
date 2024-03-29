package events

import (
	"errors"
	"sync"
)

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")
var ErrHandlerNotFound = errors.New("handler not found")

type eventDispatcher struct {
	handlers map[string][]EventHandler
}

var _ EventDispatcher = (*eventDispatcher)(nil)

func NewEventDispatcher() EventDispatcher {
	return &eventDispatcher{
		handlers: make(map[string][]EventHandler),
	}
}

func (ed *eventDispatcher) Register(eventName string, handler EventHandler) error {
	if _, ok := ed.handlers[eventName]; !ok {
		ed.handlers[eventName] = make([]EventHandler, 0)
	} else {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}

	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}

func (ed *eventDispatcher) Unregister(eventName string, handler EventHandler) error {
	if handlers, ok := ed.handlers[eventName]; ok {
		for i, h := range handlers {
			if h == handler {
				ed.handlers[eventName] = append(handlers[:i], handlers[i+1:]...)
				return nil
			}
		}
	}

	return ErrHandlerNotFound
}

func (ed *eventDispatcher) Has(eventName string, handler EventHandler) bool {
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return true
			}
		}
	}
	return false
}

func (ed *eventDispatcher) Dispatch(event Event) error {
	eventName := event.GetName()
	if handlers, ok := ed.handlers[eventName]; ok {
		wg := &sync.WaitGroup{}
		for _, handler := range handlers {
			wg.Add(1)
			go handler.Handle(event, wg)
		}
		wg.Wait()
	}

	return nil
}

func (ed *eventDispatcher) Clear() {
	ed.handlers = make(map[string][]EventHandler)
}
