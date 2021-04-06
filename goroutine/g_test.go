package goroutine

import (
	"testing"
	"time"
)

// 说明子goroutine如果没有recover()，会导致整个进程panic
func TestPanic(t *testing.T) {
	go func() {
		a, b := 0, 1
		_ = b / a
	}()
	time.Sleep(1 * time.Second)
}

func TestNotPanic(t *testing.T) {
	go func() {
		defer func() {
			err := recover()
			t.Log(err)
		}()
		a, b := 0, 1
		_ = b / a
	}()
	time.Sleep(1 * time.Second)
}
