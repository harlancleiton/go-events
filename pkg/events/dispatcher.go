package events

import "errors"

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")
var ErrHandlerNotRegistered = errors.New("handler not registered")
var ErrHandlerNotFound = errors.New("handler not found")

type EventDispatcher interface {
	Dispatch(event Event) error
	RegisterHandler(eventName string, handler EventHandler) error
	UnregisterHandler(eventName string, handler EventHandler) error
	Has(eventName string, handler EventHandler) bool
	Clear()
}

type ConcreteEventDispatcher struct {
	handlers map[string][]EventHandler
}

func NewConcreteEventDispatcher() *ConcreteEventDispatcher {
	return &ConcreteEventDispatcher{
		handlers: make(map[string][]EventHandler),
	}
}

func (d *ConcreteEventDispatcher) RegisterHandler(eventName string, handler EventHandler) error {
	if _, ok := d.handlers[eventName]; ok {
		for _, h := range d.handlers[eventName] {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}
	d.handlers[eventName] = append(d.handlers[eventName], handler)
	return nil
}

func (d *ConcreteEventDispatcher) UnregisterHandler(eventName string, handler EventHandler) error {
	if _, ok := d.handlers[eventName]; !ok {
		return ErrHandlerNotRegistered
	}

	for i, h := range d.handlers[eventName] {
		if h == handler {
			d.handlers[eventName] = append(d.handlers[eventName][:i], d.handlers[eventName][i+1:]...)
			return nil
		}
	}

	return ErrHandlerNotRegistered
}

func (d *ConcreteEventDispatcher) Dispatch(event Event) error {
	if _, ok := d.handlers[event.Name()]; !ok {
		return ErrHandlerNotFound
	}

	for _, handler := range d.handlers[event.Name()] {
		handler.Handle(event)
	}
	return nil
}

func (d *ConcreteEventDispatcher) Has(eventName string, handler EventHandler) bool {
	if _, ok := d.handlers[eventName]; !ok {
		return false
	}

	for _, h := range d.handlers[eventName] {
		if h == handler {
			return true
		}
	}
	return false
}

func (d *ConcreteEventDispatcher) Clear() {
	d.handlers = make(map[string][]EventHandler)
}
