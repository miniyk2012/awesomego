package main

import (
	"fmt"

	"golang.org/x/exp/slices"

	"github.com/miniyk2012/awesomego/open_course/mit_algorithm/homework1/tools"
)

// InversionCountSort 对arr排序并返回逆序数
func InversionCountSort(arr []int) int {
	if len(arr) <= 1 {
		return 0
	}
	if len(arr) == 2 {
		if arr[0] > arr[1] {
			arr[1], arr[0] = arr[0], arr[1]
			return 1
		} else {
			return 0
		}
	}
	left, right := arr[:len(arr)/2], arr[len(arr)/2:]
	a1 := InversionCountSort(left)
	a2 := InversionCountSort(right)
	a3 := splitCountSort(left, right, arr)
	return a1 + a2 + a3
}

// 计算分别在left,right的逆序数对个数, 并且排序到arr中
func splitCountSort(left []int, right []int, arr []int) int {
	tmp := make([]int, len(arr))
	count, i, j := 0, 0, 0
	for k := 0; k < len(arr); k++ {
		if i >= len(left) {
			tmp[k] = right[j]
			j++
		} else if j >= len(right) {
			tmp[k] = left[i]
			i++
		} else if left[i] > right[j] { // 当右边放进去1个数, 左边剩下数的都与右边该数构成逆序对
			count += len(left) - i
			tmp[k] = right[j]
			j++
		} else {
			tmp[k] = left[i]
			i++
		}
	}
	copy(arr, tmp)
	return count
}

func SlowInversionCount(arr []int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				count++
			}
		}
	}
	return count
}

func main() {
	arr := tools.ReadNumbers("./open_course/mit_algorithm/homework1/week2/IntegerArray.txt")
	copyArr := slices.Clone(arr)
	fmt.Printf("len=%d\n", len(arr))
	num := InversionCountSort(arr)
	fmt.Printf("inversionCount=%d\n", num)
	fmt.Printf("is sorted? %t\n", slices.IsSorted(arr))
	fmt.Println(arr[:10])
	fmt.Println(arr[len(arr)-10 : len(arr)])
	fmt.Printf("slow method=%d", SlowInversionCount(copyArr))
}
