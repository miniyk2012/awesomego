package main

func nextPermutation(nums []int) {
	var (
		left, right = -1, -1
	)
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			left = i - 1
			break
		}
	}
	if left == -1 {
		reverse(nums, 0, len(nums)-1)
		return
	}
	for i := len(nums) - 1; i > left; i-- {
		if nums[i] > nums[left] {
			right = i
			break
		}
	}
	nums[left], nums[right] = nums[right], nums[left]
	reverse(nums, left+1, len(nums)-1)
}

func reverse(nums []int, left int, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
