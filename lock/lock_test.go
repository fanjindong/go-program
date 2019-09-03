package lock

import (
	"sync"
	"testing"
	"time"

	"gotest.tools/assert"
)

var count int

func TestMutex(t *testing.T) {
	N := 10000
	for i := 0; i < N; i++ {
		go func() {
			count++
		}()
	}
	time.Sleep(1 * time.Second)
	assert.Assert(t, count != N, count)

	mutex := &sync.Mutex{}
	count = 0
	for i := 0; i < N; i++ {
		go func() {
			mutex.Lock()
			count++
			mutex.Unlock()
		}()
	}
	time.Sleep(1 * time.Second)
	assert.Assert(t, count == N, count)
}
