package runtime

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func deadloop() {
	for {
	}
}

func TestDeadLoop(t *testing.T) {
	runtime.GOMAXPROCS(1)
	go deadloop()
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("I am main")
	}
}
