// map法
package main

import "sort"

func threeSum(nums []int) [][]int {
	var ret [][]int = make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break // 排序后第一个都大于0, 肯定不可能和为0
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue // 为了不重复
		}
		first := nums[i]
		visited := make(map[int]int) // value是第一次key出现的下标
		for j := i + 1; j < len(nums); j++ {
			if idx, ok := visited[-nums[j]-first]; ok {
				if j > i+1 && nums[j] == nums[j-1] && idx != j-1 {
					continue // 为了不重复
				}
				ret = append(ret, []int{first, -nums[j] - first, nums[j]})
			}
			if _, ok := visited[nums[j]]; !ok { // 为了不覆盖第一次出现的下标
				visited[nums[j]] = j
			}
		}
	}
	return ret
}
