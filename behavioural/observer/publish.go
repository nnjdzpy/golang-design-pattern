package observer

import (
	"reflect"
	"sync"
)

type Publisher struct {
	Observers []Observer
}

func (p *Publisher) Subscribe(observer Observer) {
	p.Observers = append(p.Observers, observer)
}

func (p *Publisher) UnSubscribe(observer Observer) {
	newObservers := make([]Observer, len(p.Observers)-1)
	for _, o := range p.Observers {
		if !reflect.DeepEqual(observer, o) {
			newObservers = append(newObservers, o)
		}
	}
}

func (p *Publisher) Notify() {
	var wg sync.WaitGroup
	wg.Add(len(p.Observers))
	for _, o := range p.Observers {
		go func(wg *sync.WaitGroup, observer Observer) {
			observer.Update()
			wg.Done()
		}(&wg, o)
	}
	wg.Wait()
}
