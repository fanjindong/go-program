package option

import (
	"gotest.tools/assert"
	"strconv"
	"testing"
)

func TestOption(t *testing.T) {
	option := NewOption()
	option.Default = "fjd"

	m := NewMap(option)
	assert.Equal(t, len(m), option.Max-option.Min+1)
	for _, value := range m {
		assert.Equal(t, value, option.Default)
	}
}

type Option struct {
	Default string
	Min     int
	Max     int
}

func NewOption() *Option {
	return &Option{
		Default: "10098",
		Min:     10,
		Max:     20,
	}
}

func NewMap(option *Option) map[string]interface{} {
	m := make(map[string]interface{}, option.Max-option.Min)
	for i := option.Min; i <= option.Max; i++ {
		m[strconv.Itoa(i)] = option.Default
	}
	return m
}
