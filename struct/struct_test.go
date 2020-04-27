package _struct

import "testing"

type Base struct {
	A int
	B int
}

func (b Base) Sum() int {
	return b.A + b.B
}

func (b *Base) Swap() {
	b.A, b.B = b.B, b.A
}

func TestSubStruct(t *testing.T) {
	type One struct {
		Base
	}
	type Two struct {
		*Base
	}
	one := One{Base{A: 1, B: 2}}
	t.Log(one.Sum())
	one.Swap()
	t.Log(one, one.A, one.B)

	two := Two{&Base{A: 1, B: 2}}
	t.Log(two.Sum())
	two.Swap()
	t.Log(two, two.A, two.B)
}
