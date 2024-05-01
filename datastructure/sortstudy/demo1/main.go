package main

import (
	"fmt"
	"sort"
)

// StuScore 学生成绩结构体
type StuScore struct {
	name  string // 姓名
	score int    // 成绩
}

type StuScores []StuScore

func (s StuScores) Len() int {
	return len(s)
}

// Less . 排完序后, i在j前<=>Less(i,j)=true
// 也可以这么思考: 要想i排在j前, 必须满足Less(i,j)=true
func (s StuScores) Less(i, j int) bool {
	return s[i].score < s[j].score
}

func (s StuScores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func sort1() {
	stus := StuScores{
		{"alan", 95},
		{"hikerell", 91},
		{"acmfly", 96},
		{"leao", 90},
	}
	fmt.Printf("before sort, isSorted=%t\n", sort.IsSorted(stus))
	sort.Stable(stus)
	fmt.Printf("after sort, isSorted=%t\n", sort.IsSorted(stus))
	fmt.Println("Sorted:\n\t", stus)

	sort.Stable(sort.Reverse(stus))
	fmt.Println("Sorted:\n\t", stus)
	fmt.Printf("after reversed sort, isSorted=%t\n", sort.IsSorted(sort.Reverse(stus)))
}

func binary() {
	x := 11
	s := []int{3, 6, 8, 11, 45} // 注意已经升序排序
	pos := sort.Search(len(s), func(i int) bool {
		return s[i] >= x
	})
	if pos < len(s) && s[pos] == x {
		fmt.Println(x, " 在 s 中的位置为：", pos)
	} else {
		fmt.Println(x, " 在 s 中不存在")
	}
}

func GuessingGame() {
	var s string
	fmt.Printf("Pick an integer from 0 to 100.\n")
	answer := sort.Search(100, func(i int) bool {
		fmt.Printf("Is your number <= %d? ", i)
		fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)
}

func main() {
	GuessingGame()
}
