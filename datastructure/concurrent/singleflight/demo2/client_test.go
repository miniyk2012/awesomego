package main

import (
	"sync"
	"testing"
	"time"
)

func TestGetToken(t *testing.T) {
	numTasks := 1000

	// 1 business, random region
	business := "gofromzero"
	regions := []string{RegionAmerica, RegionEurope, RegionAsia, RegionAfrica}

	// run multiple tasks
	t.Logf("start %d tasks...", numTasks)
	var wg sync.WaitGroup
	for i := 0; i < numTasks; i++ {
		num := i + 1
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			idx := n % len(regions)
			region := regions[idx]
			task := newGetTokenTask(region, business)
			token := task.Do()
			if token == "" {
				t.Logf("[task:%d] get token failed!", n)
			} else {
				t.Logf("[task:%d] got token -> %s", n, token)
			}
		}(num)
		time.Sleep(1 * time.Millisecond)
	}
	wg.Wait()

	t.Logf("finish!")
}
