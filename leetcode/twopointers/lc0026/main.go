package main

import "fmt"

func removeDuplicates(nums []int) int {
	idx := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			idx++
			nums[idx] = nums[i]
		}
	}
	return idx + 1
}

func main() {
	var nums []int = []int{1, 1, 2, 3, 4, 4, 5}
	num := removeDuplicates(nums)
	fmt.Printf("%d %v\n", num, nums)
}
