package main

func removeDuplicates1(nums []int) int {
	left, right := 0, 0 // 下一步要访问的位置
	repeated := 0       // 当前重复的个数
	for ; right < len(nums); right++ {
		if repeated == 2 && nums[right] != nums[left-1] {
			nums[left] = nums[right]
			left++
			repeated = 1
		} else if repeated == 0 {
			left++
			repeated = 1
		} else if repeated == 1 && nums[right] == nums[left-1] {
			nums[left] = nums[right]
			left++
			repeated++
		} else if repeated == 1 {
			nums[left] = nums[right]
			left++
		}
	}
	return left
}

// 考虑到: nums是递增的, 因此无须记录重复个数, 只要判断新的值和2个位置之前的值是否相同
func removeDuplicates(nums []int) int {
	slow := 0 // 指向去重数组的下一个位置
	for fast, v := range nums {
		if fast < 2 || nums[slow-2] != v {
			nums[slow] = v
			slow++
		}
	}
	return slow
}
