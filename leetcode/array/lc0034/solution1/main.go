package main

func searchRange(nums []int, target int) []int {
	i, j := -1, -1
	start := false
	for idx, val := range nums {
		if val == target && !start {
			start = true
			i = idx
		}
		if val == target {
			j = idx
		} else if start {
			break // 提早跳出
		}
	}
	return []int{i, j}
}
