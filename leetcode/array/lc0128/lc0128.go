package main

import "sort"

func longestConsecutive1(nums []int) int {
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}
	longest := 0
	for num := range numSet {
		// 看num开始有多长
		if numSet[num-1] {
			continue
		}
		i := 1
		for {
			if !numSet[num+i] {
				if i > longest {
					longest = i
				}
				break
			}
			i++
		}
	}
	return longest
}

func longestConsecutive2(nums []int) int {
	// 排序不满足计算复杂度, 不过是一种解法
	if len(nums) < 2 {
		return len(nums)
	}
	sort.Ints(nums)
	longest := 0
	currentLength := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1]+1 {
			currentLength += nums[i] - nums[i-1]
		} else {
			if currentLength > longest {
				longest = currentLength
			}
			currentLength = 1
		}
	}
	if currentLength > longest {
		longest = currentLength
	}
	return longest
}
