package goevent

import (
	"reflect"
)

type EventBus struct {
	Subscribers []*Subscriber
}

func NewEventBus() *EventBus {
	return &EventBus{[]*Subscriber{}}
}

func (this *EventBus) Subscribe(subscriber interface{}) error {
	_subscriber, err := NewSubscriber(subscriber)

	if err != nil {
		return err
	}

	this.Subscribers = append(this.Subscribers, _subscriber)

	return nil
}

func (this *EventBus) Publish(event interface{}) {
	name := reflect.ValueOf(event).Type().String()
	for _, subscriber := range this.Subscribers {
		if subscriber.subscribeTo == name || subscriber.subscribeTo == "interface {}" {
			params := []reflect.Value{}
			params = append(params, reflect.ValueOf(event))
			reflect.ValueOf(subscriber.subscriber).Call(params)
		}
	}
}
