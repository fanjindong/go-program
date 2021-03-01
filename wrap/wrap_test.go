package wrap

import (
	"fmt"
	"testing"
	"time"
)

func Sum(x, y int) int {
	return x + y
}

type SumFunc func(x, y int) int

func wrapCostTime(sum SumFunc) SumFunc {
	return func(x, y int) int {
		start := time.Now()
		defer fmt.Println("cost: ", time.Since(start))
		result := sum(x, y)
		return result
	}
}

func TestWrapSum(t *testing.T) {
	sumWrap := wrapCostTime(Sum)
	result := sumWrap(1, 2)
	t.Log(result)
}

type S struct {
	A func()
}

func NewS() *S {
	s := &S{}
	s.A = func() { s.A() }
	return s
}

func (s *S) A() {
}
