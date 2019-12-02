package iota

import (
	"gotest.tools/assert"
	"testing"
)

const (
	Man int = iota
	Woman
)

func TestSex(t *testing.T) {
	assert.Equal(t, Man, 0)
	assert.Equal(t, Woman, 1)
}

const (
	On  int = 1 << iota // 1 << 0 which is 00000001
	Off                 // 1 << 1 which is 00000010
	Set                 // 1 << 2 which is 00000100
)

func TestSwitch(t *testing.T) {
	assert.Equal(t, On, 1)
	assert.Equal(t, Off, 2)
	assert.Equal(t, Set, 4)
	t.Log(On&Off, On|Off)
}
