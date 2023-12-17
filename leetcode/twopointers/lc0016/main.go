package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var ret int
	minDiff := math.MaxInt
	for i := 0; i < len(nums); i++ {
		remain := target - nums[i]
		l, r := i+1, len(nums)-1
		for l < r {
			diff := remain - nums[l] - nums[r]
			if diff == 0 {
				return target
			} else if diff < 0 {
				r--
				if -diff < minDiff {
					ret = target - diff
					minDiff = -diff
				}
			} else {
				l++
				if diff < minDiff {
					ret = target - diff
					minDiff = diff
				}
			}
		}
	}
	return ret
}
