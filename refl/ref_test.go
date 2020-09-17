package refl

import (
	"reflect"
	"testing"
)

type S struct {
	A string
	B int32
}

func TestRefl(t *testing.T) {
	s := &S{
		A: "a",
		B: 1,
	}
	v := reflect.ValueOf(s)
	ve := v.Elem()
	t.Log(ve.String())
	t.Log(v.Elem().CanSet())
	t.Log(v.CanSet())
}
