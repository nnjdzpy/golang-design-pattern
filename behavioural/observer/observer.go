package observer

import "fmt"

type Observer interface {
	Update()
}

type ConcreteObserverA struct {
	Name string
}

var _ Observer = (*ConcreteObserverA)(nil)

func (c *ConcreteObserverA) Update() {
	fmt.Print("ObserverA update")
}

type ConcreteObserverB struct {
	Name string
}

var _ Observer = (*ConcreteObserverB)(nil)

func (c *ConcreteObserverB) Update() {
	fmt.Print("ObserverB update")
}

type ConcreteObserverC struct {
	Name string
}

var _ Observer = (*ConcreteObserverC)(nil)

func (c *ConcreteObserverC) Update() {
	fmt.Print("ObserverC update")
}
