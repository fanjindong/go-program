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

func TestRWMutex(t *testing.T) {
	count = 0
	var assertCount int
	N := 1000
	mutex := &sync.RWMutex{}
	for i := 0; i < N; i++ {
		go func() {
			mutex.RLock()
			assert.Assert(t, count == assertCount, count)
			t.Log(count)
			mutex.RUnlock()
		}()
	}
	N = N / 100
	for i := 0; i < N; i++ {
		go func() {
			mutex.Lock()
			count++
			assertCount++
			t.Log(count, assertCount)
			mutex.Unlock()
		}()
	}
	time.Sleep(1 * time.Second)
	assert.Assert(t, count == N, count)
}
