package go_cache

import (
	"github.com/fanjindong/go-cache"
	pcache "github.com/patrickmn/go-cache"
	"runtime"
	"strconv"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	runtime.SetCPUProfileRate(500)
	c := cache.NewMemCache()
	for i := 0; i < 10000; i++ {
		c.Set(strconv.FormatInt(int64(i), 10), i, cache.WithEx(time.Duration(i)*time.Millisecond))
	}
	time.Sleep(3 * time.Second)
	//for i := 0; i < 10000; i++ {
	//	if _, ok := c.Get(strconv.FormatInt(int64(i), 10)); ok {
	//		//t.Log(i, "exist")
	//		//break
	//	}
	//}
}

func TestPatrickmn(t *testing.T) {
	runtime.SetCPUProfileRate(500)
	c := pcache.New(10*time.Minute, 1*time.Second)
	for i := 0; i < 10000; i++ {
		c.Set(strconv.FormatInt(int64(i), 10), i, time.Duration(i)*time.Millisecond)
	}
	time.Sleep(3 * time.Second)
	//for i := 0; i < 10000; i++ {
	//	if _, ok := c.Get(strconv.FormatInt(int64(i), 10)); ok {
	//		//t.Log(i, "exist")
	//		//break
	//	}
	//}
}
