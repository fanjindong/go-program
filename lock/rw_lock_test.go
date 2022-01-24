package lock

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var rw = sync.RWMutex{}

func rlock(v interface{}) {
	fmt.Println("read", v, "RLock")
	rw.RLock()
	time.Sleep(1 * time.Second)
	rw.RUnlock()
	fmt.Println("read", v, "RUnlock")
}

func wlock(v interface{}) {
	fmt.Println("write", v, "Lock")
	rw.Lock()
	time.Sleep(1 * time.Second)
	rw.Unlock()
	fmt.Println("write", v, "Unlock")
}

func TestRWMutexBlock(t *testing.T) {
	for i := 0; i < 10; i++ {
		go rlock(i)
	}
	go wlock(11)
	for i := 12; i < 20; i++ {
		go rlock(i)
	}
	time.Sleep(50 * time.Second)
}

func TestRWMutexRBlock(t *testing.T) {
	for i := 0; i < 10; i++ {
		go rlock(i)
	}
	go wlock(11)
	time.Sleep(100 * time.Millisecond)
	for i := 12; i < 20; i++ {
		go rlock(i)
	}
	time.Sleep(50 * time.Second)
}
