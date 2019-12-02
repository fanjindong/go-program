package option

import "testing"

type Mapper map[string]interface{}

func New(key string, opts ...Option) Mapper {
	m := Mapper{key: ""}
	for _, opt := range opts {
		opt(&m)
	}
	return m
}

type Option func(m *Mapper)

func WithDefault(value interface{}) Option {
	return func(m *Mapper) {
		for key, _ := range *m {
			(*m)[key] = value
		}
	}
}

func TestOption(t *testing.T) {
	m := New("name", WithDefault("fjd"))
	t.Log(m)
}
