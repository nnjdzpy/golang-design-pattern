package chainofresponsibility

type Handler interface {
	Handle()
	SetNext(handler Handler)
}

type ConcreteHandlerA struct {
	Next Handler
}

var _ Handler = (*ConcreteHandlerA)(nil)

func (c *ConcreteHandlerA) Handle() {
	//TODO implement me
	panic("implement me")
}

func (c *ConcreteHandlerA) SetNext(handler Handler) {
	//TODO implement me
	panic("implement me")
}

type ConcreteHandlerB struct {
	Next Handler
}

var _ Handler = (*ConcreteHandlerB)(nil)

func (c *ConcreteHandlerB) Handle() {
	//TODO implement me
	panic("implement me")
}

func (c *ConcreteHandlerB) SetNext(handler Handler) {
	//TODO implement me
	panic("implement me")
}
