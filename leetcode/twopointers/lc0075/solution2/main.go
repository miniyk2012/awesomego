package main

func sortColors(nums []int) {
	count0 := swapColors(nums, 0)
	swapColors(nums[count0:], 1)
}

func swapColors(colors []int, target int) (countTarget int) {
	for i := 0; i < len(colors); i++ {
		if colors[i] == target {
			colors[i], colors[countTarget] = colors[countTarget], colors[i]
			countTarget++
		}
	}
	return countTarget
}
