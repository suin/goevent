package goevent

import (
	"errors"
	"fmt"
	"reflect"
)

type Subscriber struct {
	subscribeTo string
	subscriber  interface{}
}

func NewSubscriber(subscriber interface{}) (*Subscriber, error) {
	if reflect.ValueOf(subscriber).Type().Kind() != reflect.Func {
		return &Subscriber{}, errors.New(fmt.Sprintf("Subscriber is not a function: %#v", subscriber))
	}
	subscribeTo := reflect.ValueOf(subscriber).Type().In(0).String()
	return &Subscriber{subscribeTo, subscriber}, nil
}
