package proxy

import "fmt"

type Service interface {
	Handle()
}

type ServiceImpl struct {
}

func (s *ServiceImpl) Handle() {
	fmt.Print("handle")
}

type Proxy struct {
	Service *ServiceImpl
}

func (p *Proxy) Handle() {
	p.BeforeHandle()
	p.Service.Handle()
	p.AfterHandle()
}

func (p *Proxy) BeforeHandle() {
	fmt.Print("before handle")
}

func (p *Proxy) AfterHandle() {
	fmt.Print("after handle")
}
