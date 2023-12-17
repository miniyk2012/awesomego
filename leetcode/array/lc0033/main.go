package main

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if target < nums[mid] {
			if nums[mid] <= nums[right] {
				right = mid - 1
			} else {
				if target >= nums[left] {
					right = mid - 1
				} else {
					left = mid + 1
				}
			}
		} else {
			if nums[mid] >= nums[left] {
				left = mid + 1
			} else {
				if target > nums[right] {
					right = mid - 1
				} else {
					left = mid + 1
				}
			}
		}
	}
	return -1
}
