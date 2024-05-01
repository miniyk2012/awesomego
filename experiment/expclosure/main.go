package main

import (
	"fmt"
	"sync"
)

func printValue(i1, i2 int) {
	fmt.Println(i1, i2)
}

func batchExecute(tasks []func()) {
	var wg sync.WaitGroup
	for _, task := range tasks {
		// shadow
		task := task
		wg.Add(1)
		go func() {
			defer func() {
				if err := recover(); err != nil {
				}
				wg.Done()
			}()
			task()
		}()
	}
	wg.Wait()
}
func main() {
	var tasks []func()
	for i := 0; i < 10; i++ {
		x1, x2 := i, 2*i
		tasks = append(tasks, func() {
			printValue(x1, x2)
		})
	}
	batchExecute(tasks)
}
