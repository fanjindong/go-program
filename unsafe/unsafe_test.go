package unsafe

import (
	"reflect"
	"testing"
	"unsafe"
)

type align1 struct {
	a int32
	b int64
	c int8
}

type align2 struct {
	a int32
	c int8
	b int64
}

// align1,align2内字段最大的都是int64, 大小为 8bytes，对齐按机器字确定，64 位下是 8bytes，所以将按 8bytes 对齐
//align1.a 大小 4bytes, 填充 4bytes 使对齐（后边字段已对齐，所以直接填充）
//align1.b 大小 8bytes, 已对齐
//align1.c 大小 1bytes，填充 7bytes 使对齐（后边无字段，所以直接填充）
//总大小为 8+8+8=24
//align2中将c提前后，a和c总大小 4bytes,在填充 4bytes 使对齐
//总大小为 8+8=16
//所以，合理重排字段可以减少填充，使 struct 字段排列更紧密
func TestAlign(t *testing.T) {
	a1 := align1{}
	a2 := align2{}
	t.Log(unsafe.Sizeof(a1), unsafe.Sizeof(a1.a), unsafe.Sizeof(a1.b), unsafe.Sizeof(a1.c))
	t.Log(unsafe.Sizeof(a2), unsafe.Sizeof(a2.a), unsafe.Sizeof(a2.b), unsafe.Sizeof(a2.c))
	// output:
	// 24 4 8 1
	// 16 4 8 1
}

func string2bytes(s string) []byte {
	sth := (*reflect.StringHeader)(unsafe.Pointer(&s))
	slh := &reflect.SliceHeader{
		Data: sth.Data,
		Len:  sth.Len,
		Cap:  sth.Len,
	}
	return *(*[]byte)(unsafe.Pointer(slh))
}

func bytes2string(b []byte) string {
	slh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sth := &reflect.StringHeader{
		Data: slh.Data,
		Len:  slh.Len,
	}
	return *(*string)(unsafe.Pointer(sth))
}

// unsafe.Pointer 可以与任意类型的指针 互相转化，这个例子是比较经典的 string 与 []byte 互相转化
func TestStringBytesMutualTransformation(t *testing.T) {
	t.Log(string2bytes("a"))
	t.Log(bytes2string([]byte{97}))
}

type S struct {
	A uint32
	B uint64
	C uint64
	D uint64
	E struct{}
}

// 当一个结构体的最后一个字段内存占用为0时，需要填充一个对齐位，以防止赋值时，内存不连续
func TestSizePadding(t *testing.T) {
	s := S{}
	t.Log(unsafe.Sizeof(s), unsafe.Sizeof(s.A), unsafe.Sizeof(s.B), unsafe.Sizeof(s.C), unsafe.Sizeof(s.D), unsafe.Sizeof(s.E))
	// output:
	// 40 4 8 8 8 0
}
