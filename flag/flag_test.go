package flag

import (
	"testing"
)

// -- bool Value
type boolValue bool

func newBoolValue(val bool, p *bool) *boolValue {
	*p = val
	return (*boolValue)(p)
}

func Test(t *testing.T){
	f := false
	var p = &f
	t.Log(p, f, *p)
	*p = true
	t.Log(p, f, *p)
	t.Log(boolValue(f), boolValue(*p), (*boolValue)(p))
}
