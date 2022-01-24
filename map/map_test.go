package _map

import (
	"testing"
)

func TestName(t *testing.T) {
	m := make(map[string]string)
	t.Log(m)
}

func TestForDel(t *testing.T) {
	m := make(map[int]bool)
	for i := 0; i < 10000; i++ {
		m[i] = true
	}
	var count int
	for k, _ := range m {
		count += 1
		if k%330 == 0 {
			delete(m, k)
			t.Log(k)
		}
	}
	t.Log(count)
}
