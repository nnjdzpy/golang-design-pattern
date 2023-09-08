package main

type Option func(map[string]interface{})

func WithName(name string) Option {
	return func(m map[string]interface{}) {
		m["name"] = name
	}
}

func WithStatus(status string) Option {
	return func(m map[string]interface{}) {
		m["status"] = status
	}
}

func WithToken(token [16]byte) Option {
	return func(m map[string]interface{}) {
		m["token"] = token
	}
}

func NewMap(opts ...Option) map[string]interface{} {
	m := make(map[string]interface{}, 0)
	for _, o := range opts {
		o(m)
	}
	return m
}
