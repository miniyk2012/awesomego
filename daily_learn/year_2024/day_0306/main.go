package main

import (
	"fmt"
	"time"
)

func append1(arr []int) {
	arr = append(arr, 10, 11, 12)
}

func append2(arr *[]int) {
	*arr = append(*arr, 10, 11, 12, 13)
}

// ConstantDelaySchedule represents a simple recurring duty cycle, e.g. "Every 5 minutes".
// It does not support jobs more frequent than once a second.
type ConstantDelaySchedule struct {
	Delay time.Duration
}

// Every returns a crontab Schedule that activates once every duration.
// Delays of less than a second are not supported (will round up to 1 second).
// Any fields less than a Second are truncated.
func Every(duration time.Duration) ConstantDelaySchedule {
	if duration < time.Second {
		duration = time.Second
	}
	return ConstantDelaySchedule{
		Delay: duration - time.Duration(duration.Nanoseconds())%time.Second,
	}
}

// Next returns the next time this should be run.
// This rounds so that the next activation time will be on the second.
func (schedule ConstantDelaySchedule) Next(t time.Time) time.Time {
	return t.Add(schedule.Delay - time.Duration(t.Nanosecond())*time.Nanosecond)
}

func sliceDemo() {
	var b = []int{1, 2, 3, 4, 5}
	var a = b[0:2]
	fmt.Printf("b cap=%d, len=%d\na cap=%d, len=%d\n", cap(b), len(b), cap(a), len(a))
	append2(&a)
	fmt.Printf("b cap=%d, len=%d\na cap=%d, len=%d\n", cap(b), len(b), cap(a), len(a))
	fmt.Println(a, b)

	z := make([]int, 3, 10)
	x := z[0:5:8]
	fmt.Println(z, x)
	fmt.Printf("z cap=%d, len=%d\nx cap=%d, len=%d\n", cap(z), len(z), cap(x), len(x))
}

func main() {
	//s := Every(time.Second * 3600)
	//fmt.Println(s.Delay)
	//fmt.Println(s.Next(time.Now()))
	defer func() {
		time.Sleep(time.Second * 3)
		fmt.Println("aaa")
	}()
	fmt.Println("bbb")
}
