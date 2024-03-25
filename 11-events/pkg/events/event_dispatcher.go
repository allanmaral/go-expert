package events

import "errors"

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

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
		ed.handlers[eventName] = make([]EventHandler, 1)
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
	return nil
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
	return nil
}

func (ed *eventDispatcher) Clear() {
	ed.handlers = make(map[string][]EventHandler)
}
