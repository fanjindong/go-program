package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
	"time"
)

// 案例分析：
// main函数大概 63μs时，执行到 `wg.Wait()`, 此时item的值为循环列表的最后一个元素
// 而第一个goroutine启动时在78μs
// 所以导致所有goroutine拿到的item为循环列表的最后一个元素
func main() {
	f, _ := os.Create("./trace/trace.out")
	_ = trace.Start(f)
	defer trace.Stop()

	wg := sync.WaitGroup{}
	for _, item := range []string{"a", "b", "c", "d"} {
		wg.Add(1)
		go func() {
			fmt.Println(item)
			wg.Done()
		}()
		time.Sleep(50*time.Microsecond)
	}
	wg.Wait()

}
