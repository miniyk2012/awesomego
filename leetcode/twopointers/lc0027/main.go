package main

import "fmt"

func removeElement(nums []int, val int) int {
	idx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[idx] = nums[i]
			idx++
		}
	}
	return idx
}

func main() {
	var nums []int = []int{1, 2, 3, 4}
	num := removeElement(nums, 2)
	fmt.Printf("%d %v\n", num, nums)
}
