package main

func sortColors(nums []int) {
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}
	var i int
	for j := 0; j < counts[0]; i, j = i+1, j+1 {
		nums[i] = 0
	}
	for j := 0; j < counts[1]; i, j = i+1, j+1 {
		nums[i] = 1
	}
	for j := 0; j < counts[2]; i, j = i+1, j+1 {
		nums[i] = 2
	}
}
