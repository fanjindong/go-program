package contexts

import (
	"context"
	"testing"
	"time"
)

// 问题：如何控制Goroutine之间的共生同死，防止调用方挂掉，而被调用方还在执行。
// 解决：基于context在多个Goroutine之间传递信号。多个 Goroutine 同时订阅 ctx.Done() 管道中的消息，一旦接收到取消信号就停止当前正在执行的工作并提前返回
func TestContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	go taskA(t, ctx, cancel)
	select {
	case <-ctx.Done():
		t.Log("over")
	}
	// time.Sleep(3 * time.Second)
}

func taskA(t *testing.T, c context.Context, cancel func()) {
	go taskB(t, c, cancel)
}

func taskB(t *testing.T, c context.Context, cancel func()) {
	go taskC(t, c, cancel)
}

func taskC(t *testing.T, c context.Context, cancel func()) {
	time.Sleep(4 * time.Second)
	t.Log("taskC over")
	// cancel()
}
