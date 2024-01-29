package concurrent

import (
	"context"
	"sync"
	"time"

	"github.com/marusama/cyclicbarrier"
)

func CyclicBarrierDemo() int {
	cnt := 0
	b := cyclicbarrier.NewWithAction(10, func() error {
		cnt++
		return nil
	})
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ { // create 10 goroutines (the same count as barrier parties)
		wg.Add(1)
		go func() {
			for j := 0; j < 5; j++ {

				// do some hard work 5 times
				time.Sleep(100 * time.Millisecond)

				err := b.Await(context.TODO()) // ..and wait for other parties on the barrier.
				// Last arrived goroutine will do the barrier action
				// and then pass all other goroutines to the next round
				if err != nil {
					panic(err)
				}
			}
			wg.Done()
		}()
	}

	wg.Wait()
	return cnt
}
