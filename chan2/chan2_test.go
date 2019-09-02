package chan2

import (
	"strconv"
	"sync"
	"testing"
	"time"
)

// 需求：goroutine 执行多个任务，如何按顺序返回结果
// 方案：任务以channel的形式按顺序注册进入一个有容量的channel, 然后阻塞的从容量channel中拿结果
func TestChan2(t *testing.T) {
	ch2 := make(chan chan string, 10)
	wg := &sync.WaitGroup{}

	for i := 0; i < 4; i++ {
		wg.Add(1)
		ch := make(chan string)
		ch2 <- ch
		go task(i, ch, wg)
	}

	go listen(ch2, t)
	wg.Wait()
}

func task(i int, ch chan string, wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	ch <- strconv.Itoa(i)
	wg.Done()
}

func listen(ch2 chan chan string, t *testing.T) {
	for ch := range ch2 {
		t.Log(<-ch)
	}
}
