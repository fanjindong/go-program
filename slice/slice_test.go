package slice

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

var _ reflect.SliceHeader

// 定理：
// 1. 切片传递是 reflect.SliceHeader 这个结构体的值传递
// 2. 值传递共享的是 reflect.SliceHeader.Data 指向的底层数据，当底层数据触发扩容时，底层数据不再共享
func TestOne(t *testing.T) {
	s := []int{1, 2, 3}
	sliceAppend(s, 4) // [1 2 3 4]
	t.Log(s)          // [1 2 3]

	array := [5]int{1, 2, 3, 4, 5}
	sliceAppend(array[0:2], 0) // [1 2 0]
	t.Log(array)               // [1 2 0 4 5]

	setHeaderLen(s) // [1 2 3 0]
	t.Log(s)        // [1 2 3]

	t.Log(array[0:0]) // []
}

func sliceAppend(s []int, a int) {
	s = append(s, a)
	fmt.Println(s)
}

func setHeaderLen(s []int) {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	sh.Len += 1
	fmt.Println(*(*[]int)(unsafe.Pointer(sh)))
}

// 42.0 ns/op
func BenchmarkMakeSliceWithLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 4)
		s = append(s, 1)
		s = append(s, 2)
		s = append(s, 3)
		s = append(s, 4)
	}
}

// 91.9 ns/op
func BenchmarkMakeSliceWithoutLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 0)
		s = append(s, 1)
		s = append(s, 2)
		s = append(s, 3)
		s = append(s, 4)
	}
}

type S struct {
	A string
	B string
	C string
	D string
	E string
	F string
	G string
}

func TestAppend(t *testing.T) {
	s1 := make([]int, 1, 2)
	s1[0] = 1
	s2 := append(s1, 2)
	t.Log(s1, s2)

	s3 := append(s2, 3)
	t.Log(s1, s2, s3)
}

func setL(ss []S) {
	for i := range ss {
		item := ss[i]
		item.A = "l"
		ss[i] = item
	}
}
func TestA(t *testing.T) {
	var ss = []S{{A: "a"}, {A: "b"}, {A: "c"}}
	setL(ss)
	t.Log(ss)
}
