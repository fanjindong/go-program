package _struct

import (
	"fmt"
	"testing"
)

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

type A struct {
	data string
}

func (a A) Print() {
	fmt.Println("A: ", a.data)
}

type B struct {
	A
	data string
}

func (b B) Print() {
	fmt.Println("B: ", b.data)
}

// 这个例子说明，go只有结构体组合，没有继承
func TestStructCombination(t *testing.T) {
	a := A{data: "aaa"}
	b := B{data: "bbb", A: a}
	a.Print()
	b.Print()
	b.A.Print()
	// output:
	//A:  aaa
	//B:  bbb
	//A:  aaa
}
