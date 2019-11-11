package redis_test

import (
	"testing"

	"github.com/garyburd/redigo/redis"
)

func TestRedis(t *testing.T) {
	t.Log("start")
	c, err := redis.DialURL("redis://192.168.3.3:6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	c.Do("SET", "hello", "world")
	s, err := redis.String(c.Do("GET", "hello"))
	t.Log(s, err)
}
