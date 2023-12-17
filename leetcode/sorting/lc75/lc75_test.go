package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
)

func TestSortColors(t *testing.T) {
	testCases := []struct {
		nums   []int
		sorted []int
	}{
		{[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{[]int{1, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 1, 2}},
	}
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%v", testCase.nums), func(t *testing.T) {
			input := slices.Clone(testCase.nums)
			sortColors3(input)
			assert.Equal(t, input, testCase.sorted)
		})
	}
}
