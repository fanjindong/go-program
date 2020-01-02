package scheduler_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/fanjindong/scheduler"
)

func task1(t *testing.T) {
	t.Log(time.Now(), "task1")
}

func task2() {
	fmt.Println(time.Now(), "task2")
}

func TestScheduler(t *testing.T) {
	s, err := scheduler.NewScheduler(1000)
	if err != nil {
		panic(err) // just example
	}
	t.Log(time.Now(), "start")
	// delay with 1 second, job function with arguments
	s.Delay().Second(5).Minute(1).Do(task1, t)
	time.Sleep(70 * time.Second)
	//// delay with 1 minute, job function without arguments
	//s.Delay().Minute(1).Do(task2)
	//
	//// delay with 1 hour
	//s.Delay().Hour(1).Do(task2)
	//
	//// special: execute immediately
	//s.Delay().Do(task2)
	//
	//// cancel job
	//jobID := s.Delay().Day(1).Do(task2)
	//err = s.CancelJob(jobID)
	//if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println("cancel delay job success")
	//}
}
