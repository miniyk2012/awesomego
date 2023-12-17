package main

import (
	"fmt"
)

func maxArea(height []int) int {
	l, r, maxValue := 0, len(height)-1, 0
	for l < r {
		width := r - l
		var i, j, high int
		if height[l] < height[r] {
			high = height[l] // 把本次的值存下来, 然后找到下次的位置
			for i = l + 1; i < r && height[i] <= height[l]; i++ {
			}
			l = i
		} else {
			high = height[r]
			for j = r - 1; j > l && height[j] <= height[r]; j-- {
			}
			r = j
		}
		temp := width * high
		if temp > maxValue {
			maxValue = temp
		}
	}
	return maxValue
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}
