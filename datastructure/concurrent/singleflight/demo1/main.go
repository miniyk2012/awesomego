package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"

	"golang.org/x/sync/singleflight"
)

type Result string

func find(ctx context.Context, i int, query string) (Result, error) {
	time.Sleep(time.Millisecond * time.Duration(rand.Int31n(70)))
	fmt.Printf("%d wawa\n", i)
	return Result(fmt.Sprintf("%d result for %q", i, query)), nil
}

func DemoDo() {
	var g singleflight.Group
	const n = 5
	waited := int32(n)
	done := make(chan struct{})
	key := "https://weibo.com/1227368500/H3GIgngon"
	for i := 0; i < n; i++ {
		time.Sleep(time.Millisecond * time.Duration(rand.Int31n(100)))
		go func(j int) {
			v, _, shared := g.Do(key, func() (interface{}, error) {
				ret, err := find(context.Background(), j, key)
				return ret, err
			})
			if atomic.AddInt32(&waited, -1) == 0 {
				close(done)
			}
			fmt.Printf("index: %d, val: %v, shared: %v\n", j, v, shared)
		}(i)
	}

	select {
	case <-done:
		time.Sleep(time.Second)
	case <-time.After(time.Second):
		fmt.Println("Do hangs")
	}

	v, _, shared := g.Do(key, func() (interface{}, error) {
		ret, err := find(context.Background(), -1, key)
		return ret, err
	})
	fmt.Printf("index: %d, val: %v, shared: %v\n", -1, v, shared)
}

func DemoChan() {
	var g singleflight.Group
	const n = 5
	waited := int32(n)
	done := make(chan struct{})
	key := "https://weibo.com/1227368500/H3GIgngon"
	for i := 0; i < n; i++ {
		time.Sleep(time.Millisecond * time.Duration(rand.Int31n(100)))
		go func(j int) {
			ch := g.DoChan(key, func() (interface{}, error) {
				ret, err := find(context.Background(), j, key)
				return ret, err
			})
			// Create our timeout
			timeout := time.After(50 * time.Millisecond)
			var ret singleflight.Result
			select {
			case <-timeout: // Timeout elapsed
				fmt.Printf("%d Timeout\n", j)
				if atomic.AddInt32(&waited, -1) == 0 {
					close(done)
				}
				return
			case ret = <-ch: // Received result from channel
				fmt.Printf("index: %d, val: %v, shared: %v\n", j, ret.Val, ret.Shared)
				if atomic.AddInt32(&waited, -1) == 0 {
					close(done)
				}
			}

		}(i)
	}

	select {
	case <-done:
		time.Sleep(time.Second)
	case <-time.After(time.Second):
		fmt.Println("Do hangs")
	}
}

func main() {
	DemoDo()
	fmt.Printf("------------------- DemoChan -------------------")
	DemoChan()
}
