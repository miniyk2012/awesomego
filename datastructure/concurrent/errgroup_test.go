package concurrent_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"golang.org/x/sync/errgroup"
)

func TimedTask(i int, ctx context.Context, earlyReturn bool) error {
	// Do some work...
	// Use ctx to handle cancellation
	if earlyReturn && ctx.Err() != nil {
		return ctx.Err()
	}
	time.Sleep(time.Millisecond * 100)
	fmt.Printf("%d,", i)
	// Return an error to stop other goroutines
	if i == 100 {
		return fmt.Errorf("error from goroutine %d", i)
	}
	return nil
}
func TestErrGroup(t *testing.T) {
	start := time.Now()
	g, ctx := errgroup.WithContext(context.Background())
	g.SetLimit(10)

	// Launch multiple goroutines in the group
	for i := 0; i < 1000; i++ {
		i := i // capture loop variable
		// 全都进来了, 不过后面的基本都提前退出了
		g.Go(func() error {
			return TimedTask(i, ctx, true) // 10个一批, 10批是要有sleep的, 其他都不sleep, 那就大概花1s
		})
	}

	// Wait for all goroutines in the group to finish
	if err := g.Wait(); err != nil {
		fmt.Println("Received error:", err)
	}
	t.Logf("early return time cost = %v", time.Now().Sub(start))

	start = time.Now()
	for i := 0; i < 200; i++ {
		i := i // capture loop variable
		g.Go(func() error {
			return TimedTask(i, ctx, false) // 10个一批, 一共20批, 那就是2s
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Println("Received error:", err)
	}
	t.Logf("total run time cost = %v", time.Now().Sub(start))
}
