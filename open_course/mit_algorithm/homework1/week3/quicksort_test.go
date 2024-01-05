package main

import (
	"sort"
	"testing"

	"golang.org/x/exp/slices"

	"github.com/miniyk2012/awesomego/open_course/mit_algorithm/homework1/tools"
	"github.com/stretchr/testify/assert"
)

func getAssignmentInput() []int {
	return tools.ReadNumbers("./quicksort_assignment.txt")
}

func TestPickQuickSort(t *testing.T) {
	arr := getAssignmentInput()

	copyArr := slices.Clone(arr)
	v := QuickSort(copyArr, PickLeft)
	t.Log(v) // 162085
	assert.Equal(t, 162085, v)
	assert.True(t, sort.IntsAreSorted(copyArr))

	copyArr = slices.Clone(arr)
	v = QuickSort(copyArr, PickRight)
	t.Log(v) // 164123
	assert.Equal(t, 164123, v)
	assert.True(t, sort.IntsAreSorted(copyArr))

	copyArr = slices.Clone(arr)
	v = QuickSort(copyArr, PickMedianThree)
	t.Log(v) // 138382
	assert.Equal(t, 138382, v)
	assert.True(t, sort.IntsAreSorted(copyArr))

	copyArr = slices.Clone(arr)
	v = QuickSort(copyArr, PickMedianOfThree)
	t.Log(v) // 138382
	assert.Equal(t, 138382, v)
	assert.True(t, sort.IntsAreSorted(copyArr))
}
