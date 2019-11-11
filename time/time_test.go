package time_test

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	now := time.Now()
	t.Log(time.Date(now.Year()-18, now.Month(), now.Day(), 0, 0, 0, 0, time.Local).Format("2006-01-02"))
}
