package goevent

import (
	"fmt"
	. "github.com/r7kamura/gospel"
	"testing"
)

func TestDescribe(t *testing.T) {

	stringSubscriberCalled := 0
	anEventSubscriberCalled := 0
	pointerOfAnEventSubscriberCalled := 0
	interfaceSubscriberCalled := 0

	Describe(t, "Subscribe and Publish", func() {
		var bus *EventBus
		Context("Construct EventBus", func() {
			bus = NewEventBus()
			It("EventBus should have no Subscribers", func() {
				Expect(len(bus.Subscribers)).To(Equal, 0)
			})
		})

		Context("Add a subscriber which subscribes string type event and publish a string event", func() {
			bus.Subscribe(func(event string) {
				stringSubscriberCalled += 1
			})
			bus.Publish("string")

			It("the subscriber subscribes the event once", func() {
				Expect(stringSubscriberCalled).To(Equal, 1)
			})
		})

		Context("Add a subscriber which subscribes AnEvent and publish AnEvent", func() {
			bus.Subscribe(func(event AnEvent) {
				anEventSubscriberCalled += 1
			})
			bus.Publish(AnEvent{})

			It("the subscriber subscribes the event once", func() {
				Expect(anEventSubscriberCalled).To(Equal, 1)
			})
		})

		Context("Add a subscriber which subscribes pointer type of AnEvent and publish *AnEvent", func() {
			bus.Subscribe(func(event *AnEvent) {
				pointerOfAnEventSubscriberCalled += 1
			})
			bus.Publish(&AnEvent{})

			It("the subscriber subscribes the event once", func() {
				Expect(pointerOfAnEventSubscriberCalled).To(Equal, 1)
			})

			It("the subscriber which subscribes AnEvent does not subscribes the event", func() {
				Expect(anEventSubscriberCalled).To(Equal, 1)
			})
		})

		Context("Add a subscriber which subscribes `interface{}` and publish *AnEvent", func() {
			bus.Subscribe(func(event interface{}) {
				interfaceSubscriberCalled += 1
			})
			bus.Publish(&AnEvent{})

			It("interfaceSubscriber subscribes *AnEvent once", func() {
				Expect(interfaceSubscriberCalled).To(Equal, 1)
			})
		})

		Context("Add a subscriber which is not a function type", func() {
			err := bus.Subscribe("string")

			It("an error occurres", func() {
				Expect(err).To(NotEqual, nil)
				Expect(fmt.Sprintf("%s", err)).To(Equal, "Subscriber is not a function: \"string\"")
			})

		})

		Context("When-Then DSL", func() {

			doSomethingCalled := 0

			doSomething := func(event *SomethingDone) {
				doSomethingCalled += 1
			}

			bus.When(SomethingDone{}).Then(doSomething)
			bus.Publish(&SomethingDone{})

			It("doSomething is called", func() {
				Expect(doSomethingCalled).To(Equal, 1)

			})
		})

		Context("Event types are different", func() {
			doSomething := func(event *SomethingDone) {}
			err := bus.When(DifferentEvent{}).Then(doSomething)

			It("an error occurres", func() {
				Expect(err).To(NotEqual, nil)
				Expect(fmt.Sprintf("%s", err)).To(Equal, `Event type does not match: goevent.DifferentEvent. Subscriber expects *goevent.SomethingDone`)
			})
		})

		Context("Add a subscriber which is not a function type", func() {
			err := bus.When("string").Then("string")

			It("an error occurres", func() {
				Expect(err).To(NotEqual, nil)
				Expect(fmt.Sprintf("%s", err)).To(Equal, "Subscriber is not a function: \"string\"")
			})

		})

		Context("Purge subscribers", func() {
			bus.Purge()

			It("should be no subscribers", func() {
				Expect(len(bus.Subscribers)).To(Equal, 0)
			})
		})
	})
}

type AnEvent struct {
	Name string
}

type SomethingDone struct {
}

type DifferentEvent struct {
}
