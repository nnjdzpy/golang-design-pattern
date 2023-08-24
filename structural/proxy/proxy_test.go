package proxy

import "testing"

func Test(t *testing.T) {
	service := &ServiceImpl{}
	proxy := &Proxy{
		Service: service,
	}

	proxy.Handle()
}
