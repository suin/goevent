package goevent

import (
	"errors"
	"fmt"
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

// mainly for testing
func (this *EventBus) Purge() {
	this.Subscribers = []*Subscriber{}
}

func (this *EventBus) When(event interface{}) *When {
	eventType := reflect.ValueOf(event).Type().String()
	return &When{this, eventType}
}

type When struct {
	bus       *EventBus
	eventType string
}

func (this *When) Then(subscriber interface{}) error {
	_subscriber, err := NewSubscriber(subscriber)

	if err != nil {
		return err
	}

	if this.eventType == _subscriber.subscribeTo || "*"+this.eventType == _subscriber.subscribeTo {
		this.bus.Subscribers = append(this.bus.Subscribers, _subscriber)
		return nil

	} else {
		return errors.New(fmt.Sprintf("Event type does not match: %s. Subscriber expects %s", this.eventType, _subscriber.subscribeTo))
	}
}
