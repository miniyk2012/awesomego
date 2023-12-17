package main

import (
	"math/rand"
	"slices"
	"testing"

	"github.com/miniyk2012/awesomego/open_course/mit_algorithm/homework1/tools"
	"github.com/stretchr/testify/assert"
)

func getAssignmentInput() []int {
	return tools.ReadNumbers("./IntegerArray.txt")
}
func generateSortedArray(n int, reverse bool) []int {
	arr := make([]int, n)
	for i := 1; i <= n; i++ {
		idx := i - 1
		if reverse {
			idx = n - 1 - idx
		}
		arr[idx] = i
	}
	return arr
}

func generateRandomArray(n int) []int {
	arr := make([]int, n)
	for i := 1; i <= n; i++ {
		idx := i - 1
		arr[idx] = i
	}
	rand.Shuffle(n, func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	return arr
}

func TestUtils(t *testing.T) {
	a := generateSortedArray(4, true)
	t.Log(a)
	a = generateSortedArray(4, false)
	t.Log(a)
	a = generateRandomArray(10)
	t.Log(a)
	a = generateSortedArray(10, true)
	assert.Equal(t, 45, SlowInversionCount(a))
	a = generateSortedArray(10, false)
	assert.Equal(t, 0, SlowInversionCount(a))
}

func TestRegularInversionCount(t *testing.T) {
	// 验证顺序和倒序的情况
	a := generateSortedArray(100, true)
	copyA := slices.Clone(a)
	assert.Equal(t, SlowInversionCount(a), InversionCountSort(copyA))

	b := generateSortedArray(100, true)
	copyB := slices.Clone(b)
	assert.Equal(t, SlowInversionCount(b), InversionCountSort(copyB))
	arr := getAssignmentInput()

	assert.Equal(t, 2407905288, InversionCountSort(arr))
}

func TestRandInversionCount(t *testing.T) {
	ns := []int{1, 2, 56, 123, 345, 6764, 10790}
	for _, n := range ns {
		arr := generateRandomArray(n)
		copyArr := slices.Clone(arr)
		assert.Equal(t, SlowInversionCount(arr), InversionCountSort(copyArr))
	}
}
