package goevent

import (
    . "github.com/r7kamura/gospel"
    "testing"
    "fmt"
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
            bus.Subscribe(func(event string){
                stringSubscriberCalled += 1
            })
            bus.Publish("string")

            It("the subscriber subscribes the event once", func() {
                Expect(stringSubscriberCalled).To(Equal, 1)
            })
        })

        Context("Add a subscriber which subscribes AnEvent and publish AnEvent", func(){
            bus.Subscribe(func(event AnEvent){
                anEventSubscriberCalled += 1
            })
            bus.Publish(AnEvent{})

            It("the subscriber subscribes the event once", func() {
                Expect(anEventSubscriberCalled).To(Equal, 1)
            })
        })

        Context("Add a subscriber which subscribes pointer type of AnEvent and publish *AnEvent", func(){
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
            bus.Subscribe(func(event interface{}){
                interfaceSubscriberCalled += 1
            })
            bus.Publish(&AnEvent{})

            It("interfaceSubscriber subscribes *AnEvent once", func(){
                Expect(interfaceSubscriberCalled).To(Equal, 1)
            })
        })

        Context("Add a subscriber which is not a function type", func(){
            err := bus.Subscribe("string")

            It("an error occurres", func(){
                Expect(err).To(NotEqual, nil)
                Expect(fmt.Sprintf("%s",err)).To(Equal, "Subscriber is not a function: \"string\"")
            })

        })
    })
}


type AnEvent struct {

}