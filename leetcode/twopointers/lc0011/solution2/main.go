package main

func maxArea(height []int) int {
	start, end, maxValue := 0, len(height)-1, 0
	for start < end {
		width := end - start
		var high int
		if height[start] < height[end] {
			high = height[start]
			start++
		} else {
			high = height[end]
			end--
		}
		temp := width * high
		if temp > maxValue {
			maxValue = temp
		}
	}
	return maxValue
}
