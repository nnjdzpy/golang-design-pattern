package observer

import (
	"testing"
)

func TestPublisher_Notify(t *testing.T) {
	observerA := &ConcreteObserverA{}
	observerB := &ConcreteObserverB{}
	observerC := &ConcreteObserverC{}

	publisher := &Publisher{Observers: []Observer{}}

	publisher.Subscribe(observerA)
	publisher.Subscribe(observerB)
	publisher.Subscribe(observerC)

	publisher.Notify()
}
