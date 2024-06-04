package app2

import (
	"fmt"
	"sync"
)

var (
	setupDone bool
	setupOnce sync.Once
)

func Setup() {
	setupOnce.Do(func() {
		// 执行一次性的设置代码
		setupDone = true
		fmt.Println("setup")
	})
}

func Teardown() {
	if setupDone {
		// 执行清理代码
		fmt.Println("tear down")
		setupDone = false
	}
}
