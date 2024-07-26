package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"code.byted.org/aweme-go/dsync"
)

func main() {
	pool := dsync.NewGoPool(100000)

	var n int
	taskC := &dsync.Task{
		RunFunc: func(context.Context) error {
			n++
			fmt.Println("C", n)
			return nil
		},
		Identifier: "TaskC",
	}
	taskB := &dsync.Task{
		RunFunc: func(context.Context) error {
			n++
			fmt.Println("B", n)
			time.Sleep(time.Second)
			return nil
		},
		Identifier: "TaskB",
	}
	taskA := &dsync.Task{
		RunFunc: func(context.Context) error {
			n++
			fmt.Println("A", n)
			return nil
		},
		Identifier: "TaskA",
	}
	taskD := &dsync.Task{
		RunFunc: func(context.Context) error {
			n++
			fmt.Println("D", n)
			return errors.New("aa")
		},
		Identifier: "TaskD",
	}
	taskA.AddChildren(taskB)
	taskA.AddChildren(taskD)
	taskA.AddChildren(taskB)
	taskB.AddChildren(taskC)
	start := time.Now()
	res := pool.Do(context.Background(), taskA)
	for v := range res {
		fmt.Printf("%ds: %s\n", time.Since(start)/time.Second, v.Identifier)
	}
	// Output: A 1
	// B 2
	// C 3
}
