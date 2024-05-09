package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPivotIndex(t *testing.T) {
	testcases := []struct {
		input  []int
		output int
	}{
		{[]int{1, 7, 3, 6, 5, 6}, 3},
		{[]int{-1, -1, 0, 1, 1, 0}, 5},
		{[]int{2, 1, -1}, 0},
	}
	for _, testCase := range testcases {
		ret := pivotIndex(testCase.input)
		assert.Equal(t, ret, testCase.output)
	}
}
