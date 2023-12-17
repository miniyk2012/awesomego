package main

func sortColors(nums []int) {
	ptr := 0
	for i, num := range nums {
		// ptr-1是0的末尾
		if num == 0 {
			nums[ptr], nums[i] = nums[i], nums[ptr]
			ptr++
		}
	}
	for i := ptr; i < len(nums); i++ {
		// ptr-1是1的末尾
		if nums[i] == 1 {
			nums[ptr], nums[i] = nums[i], nums[ptr]
			ptr++
		}
	}
}

func sortColors2(nums []int) {
	// p0-1, p1-1分别是0和1的末尾
	p0, p1 := 0, 0
	for i, num := range nums {
		if num == 1 {
			nums[p1], nums[i] = nums[i], nums[p1]
			p1++
		} else if num == 0 {
			nums[p0], nums[i] = nums[i], nums[p0]
			if p1 > p0 { // 说明把1移到了i位置
				nums[p1], nums[i] = nums[i], nums[p1]
			}
			p0++
			p1++
		}
	}
}

// 见官方题解3: https://leetcode.cn/problems/sort-colors/solutions/437968/yan-se-fen-lei-by-leetcode-solution/
func sortColors3(nums []int) {
	// 循环不变量: p0-1是0的末尾, p2+1是2开头
	p0, p2 := 0, len(nums)-1
	for i := 0; i <= p2; i++ {
		for ; i <= p2 && nums[i] == 2; p2-- {
			nums[i], nums[p2] = nums[p2], nums[i]
		}
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
		}
	}
}
