package main

import "fmt"

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxArea(height []int) int {
	if len(height) == 0 {
		return 0
	}
	head, tail := 0, len(height)-1
	ret := (tail - head) * min(height[head], height[tail])
	for {
		if height[head] < height[tail] {
			head++
			if head < tail && height[head] <= height[head-1] { // 如果移动后的高度比上次还低, 则应该继续移动
				continue
			}
		} else {
			tail--
			if tail > head && height[tail] <= height[tail+1] {
				continue
			}
		}
		if head >= tail {
			break
		}
		current := (tail - head) * min(height[head], height[tail])
		if current > ret {
			ret = current
		}
	}
	return ret
}

func assert(real, expect int) {
	if real != expect {
		panic(fmt.Sprintf("%d != %d", real, expect))
	}
}
func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	assert(maxArea(height), 49)
	height = []int{1, 1}
	assert(maxArea(height), 1)
}
