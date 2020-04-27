package time_test

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	now := time.Now()
	t.Log(time.Date(now.Year()-18, now.Month(), now.Day(), 0, 0, 0, 0, time.Local).Format("2006-01-02"))
}

// func TestTimeTicker(t *testing.T) {
// 	for range time.NewTicker(3 * time.Second).C {
// 		fmt.Println(time.Now())
// 	}
// }

func TestServiceRemainSecond(t *testing.T) {
	todayLast := time.Now().Format("2006-01-02") + " 23:59:59"

	lc, _ := time.LoadLocation("Asia/Shanghai")
	todayLastTime, _ := time.ParseInLocation("2006-01-02 15:04:05", todayLast, lc)
	remainSecond := time.Duration(todayLastTime.Unix()-time.Now().Local().Unix()) * time.Second

	t.Log(remainSecond.Seconds())
	t.Log(86400 - (time.Now().Unix()-int64(1577808000))%86400)
}
