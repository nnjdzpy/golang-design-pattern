package chainofresponsibility

import (
	"context"
	"fmt"
)

type Handler interface {
	Handle(ctx context.Context)
	SetNext(handler Handler)
}

type ConcreteHandlerA struct {
	Next Handler
}

var _ Handler = (*ConcreteHandlerA)(nil)

func (c *ConcreteHandlerA) Handle(ctx context.Context) {
	fmt.Print("handlerA")
	if c.Next != nil {
		c.Next.Handle(ctx)
	}
}

func (c *ConcreteHandlerA) SetNext(handler Handler) {
	c.Next = handler
}

type ConcreteHandlerB struct {
	Next Handler
}

var _ Handler = (*ConcreteHandlerB)(nil)

func (c *ConcreteHandlerB) Handle(ctx context.Context) {
	fmt.Print("handlerB")
	if c.Next != nil {
		c.Next.Handle(ctx)
	}
}

func (c *ConcreteHandlerB) SetNext(handler Handler) {
	c.Next = handler
}
