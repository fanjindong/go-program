package p

import "testing"

func PrintMap(t *testing.T, m map[string]string) {
	t.Logf("PrintMap, 内存地址：%p", m)
}

func TestMap(t *testing.T) {
	m := make(map[string]string)
	PrintMap(t, m)
	t.Logf("mainMap, 内存地址：%p", m)
}

type S struct {
	a string
}

func PrintStruct(t *testing.T, s *S) {
	t.Logf("PrintStruct, 内存地址：%p", s)
}

func TestStruct(t *testing.T) {
	s := S{a: "123"}
	t.Logf("PrintStruct, 内存地址：%p", &s)
	PrintStruct(t, &s)
}

func PrintSlice(t *testing.T, s []string) {
	t.Logf("PrintSlice, 内存地址：%p", s)
}

func TestSlice(t *testing.T) {
	s := []string{}
	t.Logf("TestSlice, 内存地址：%p", s)
	PrintSlice(t, s)
}
