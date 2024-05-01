package main

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_waitGroup1(t *testing.T) {
	tasksNum := 100000

	dataCh := make(chan interface{}, tasksNum)
	resp := make([]interface{}, 0, tasksNum)
	// 启动读 goroutine
	go func() {
		for data := range dataCh {
			resp = append(resp, data)
		}
	}()

	// 保证获取到所有数据后，通过 channel 传递到读协程手中
	var wg sync.WaitGroup
	for i := 0; i < tasksNum; i++ {
		wg.Add(1)
		go func(ch chan<- interface{}) {
			defer wg.Done()
			ch <- time.Now().UnixNano()
		}(dataCh)
	}
	// 确保所有取数据的协程都完成了工作，才关闭 ch
	wg.Wait()
	close(dataCh)
	t.Logf("resp len: %d", len(resp))
}

func Test_waitGroup2(t *testing.T) {
	tasksNum := 100000

	dataCh := make(chan interface{})
	resp := make([]interface{}, 0, tasksNum)
	stopCh := make(chan struct{}, 1)
	// 启动读 goroutine
	go func() {
		for data := range dataCh {
			resp = append(resp, data)
		}
		stopCh <- struct{}{}
	}()

	// 保证获取到所有数据后，通过 channel 传递到读协程手中
	var wg sync.WaitGroup
	for i := 0; i < tasksNum; i++ {
		wg.Add(1)
		go func(ch chan<- interface{}) {
			defer wg.Done()
			ch <- time.Now().UnixNano()
		}(dataCh)
	}
	// 确保所有取数据的协程都完成了工作，才关闭 ch
	wg.Wait()
	close(dataCh)

	// 确保读协程处理完成
	<-stopCh

	t.Logf("resp len: %d", len(resp))
}

func Test_waitGroup3(t *testing.T) {
	tasksNum := 10000
	dataCh := make(chan any)
	// 启动写 goroutine，推进并发获取数据进程，将获取到的数据聚合到 channel 中
	go func() {
		// 保证获取到所有数据后，通过 channel 传递到读协程手中
		var wg sync.WaitGroup
		for i := 0; i < tasksNum; i++ {
			wg.Add(1)
			go func(ch chan<- any) {
				defer wg.Done()
				dataCh <- time.Now().UnixNano()
			}(dataCh)
		}
		wg.Wait()
		// 确保所有取数据的协程都完成了工作，才关闭 ch
		close(dataCh)
	}()
	resp := make([]interface{}, 0, tasksNum)
	// 主协程作为读协程，持续读取数据，直到所有写协程完成任务，chan 被关闭后才会往下
	for v := range dataCh {
		resp = append(resp, v)
	}
	t.Logf("resp len: %d", len(resp))
	assert.Equal(t, tasksNum, len(resp))
}
