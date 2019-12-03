package common

import (
	"fmt"
	"testing"
	"time"
)

type Service struct {
	t      *testing.T
	missch chan func()
}

func New(t *testing.T) *Service {
	s := &Service{missch: make(chan func(), 1024), t: t}
	go s.cacheproc() // 按序消费 chan 桶里面的函数
	return s
}

// addCache
func (s *Service) addCache(f func()) {
	select {
	case s.missch <- f:
	default:
		fmt.Println("cacheproc chan full")
	}
}

// cacheproc is a routine for executing closure.
func (s *Service) cacheproc() {
	for {
		f := <-s.missch
		f()
	}
}

func TestAddCache(t *testing.T) {
	s := New(t)
	for i := 0; i < 1000; i++ {
		value := i
		s.addCache(
			func() {
				t.Log(value)
			})
	}
	t.Log("addCache complete")
	time.Sleep(3 * time.Second)
}
