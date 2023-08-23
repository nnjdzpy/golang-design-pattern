package chainofresponsibility

import (
	"context"
	"testing"
)

func Test(t *testing.T) {

	handlerA := &ConcreteHandlerA{}
	handlerB := &ConcreteHandlerB{}
	handlerA.SetNext(handlerB)

	ctx := context.Background()
	handlerA.Handle(ctx)
}
