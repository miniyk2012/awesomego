package main

// PickFunc 表示从[left, right]区间中挑选第哪个作为pivot
type PickFunc func(arr []int, left, right int) int

// baseQuickSort 快排, 并返回比较次数
func baseQuickSort(arr []int, left, right int, pick PickFunc) int {
	if right <= left || left < 0 || right >= len(arr) {
		return 0
	}
	mid := partition(arr, left, right, pick)
	compareNum := right - left
	if mid > left {
		compareNum += baseQuickSort(arr, left, mid-1, pick)
	}
	if mid < right {
		compareNum += baseQuickSort(arr, mid+1, right, pick)
	}
	return compareNum
}

// partition 对[left, right]区间做分割, 返回分割后pivot所在位置
func partition(arr []int, left, right int, pick PickFunc) (mid int) {
	pivotPos := pick(arr, left, right)
	arr[left], arr[pivotPos] = arr[pivotPos], arr[left]
	i, j := left+1, left+1
	for j <= right {
		if arr[j] < arr[left] {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
		j++
	}
	arr[left], arr[i-1] = arr[i-1], arr[left]
	return i - 1
}

func PickLeft(arr []int, left, right int) int {
	return left
}

func PickRight(arr []int, left, right int) int {
	return right
}

func PickMedianThree(arr []int, left, right int) int {
	var mid int
	if (right-left+1)%2 == 0 {
		mid = left + (right-left+1)/2 - 1
	} else {
		mid = left + (right-left)/2
	}
	if arr[left] < arr[mid] {
		if arr[mid] < arr[right] {
			return mid
		} else {
			if arr[right] > arr[left] {
				return right
			} else {
				return left
			}
		}
	} else if arr[mid] > arr[right] {
		return mid
	} else {
		if arr[left] > arr[right] {
			return right
		} else {
			return left
		}
	}
}

func PickMedianOfThree(arr []int, left, right int) int {
	var mid int
	if (right-left+1)%2 == 0 {
		mid = left + (right-left+1)/2 - 1
	} else {
		mid = left + (right-left)/2
	}
	bigger := func(a, b int) int {
		if arr[a] > arr[b] {
			return a
		}
		return b
	}
	biggest := bigger(left, bigger(mid, right))
	if biggest == left {
		return bigger(mid, right)
	}
	if biggest == right {
		return bigger(mid, left)
	}
	return bigger(left, right)
}

func QuickSort(arr []int, pick PickFunc) int {
	return baseQuickSort(arr, 0, len(arr)-1, pick)
}
