package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var solution = longestConsecutive1

func TestLongestConsecutive(t *testing.T) {
	testcases := []struct {
		input  []int
		output int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
		{[]int{0}, 1},
		{[]int{0, 4}, 1},
	}
	for _, testCase := range testcases {
		ret := solution(testCase.input)
		assert.Equal(t, ret, testCase.output)
	}
}
