package main

import "sort"

func threeSum(nums []int) [][]int {
	ret := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			if l > i+1 && nums[l] == nums[l-1] {
				l++
				continue
			}
			if r < len(nums)-1 && nums[r] == nums[r+1] {
				r--
				continue
			}
			if nums[i]+nums[l] > 0 {
				break
			}
			total := nums[i] + nums[l] + nums[r]
			if total == 0 {
				ret = append(ret, []int{nums[i], nums[l], nums[r]})
				l++
				r--
			} else if total > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return ret
}
