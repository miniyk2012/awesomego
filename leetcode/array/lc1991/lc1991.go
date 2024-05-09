package main

// 假设和为total, 记左边和为left, 当索引遍历到第i个元素时, 右边=total-nums[i]-left, 左右两侧元素相等,
// 则说明 total-nums[i]-left == left, 因此有2*left+nums[i]==total
func pivotIndex(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	left := 0
	for idx, v := range nums {
		if 2*left+v == total {
			return idx
		}
		left += v
	}
	return -1
}
