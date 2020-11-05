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
