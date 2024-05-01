package events

import (
	"errors"
	"sync"
)

type EventDispatcher interface {
	Dispatch(event Event) error
	Register(eventName string, handler EventHandler) error
	Unregister(eventName string, handler EventHandler) error
	Has(eventName string, handler EventHandler) bool
	Clear()
}

type BaseEventDispatcher struct {
	handlers map[string][]EventHandler
}

var _ EventDispatcher = (*BaseEventDispatcher)(nil)

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")
var ErrHandlerNotFound = errors.New("handler not found")

func NewEventDispatcher() EventDispatcher {
	return &BaseEventDispatcher{
		handlers: make(map[string][]EventHandler),
	}
}

func (ed *BaseEventDispatcher) Register(eventName string, handler EventHandler) error {
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

func (ed *BaseEventDispatcher) Unregister(eventName string, handler EventHandler) error {
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

func (ed *BaseEventDispatcher) Has(eventName string, handler EventHandler) bool {
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return true
			}
		}
	}
	return false
}

func (ed *BaseEventDispatcher) Dispatch(event Event) error {
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

func (ed *BaseEventDispatcher) Clear() {
	ed.handlers = make(map[string][]EventHandler)
}
