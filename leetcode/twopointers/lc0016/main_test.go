package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestThreeSum(t *testing.T) {
	cases := []struct {
		nums     []int
		target   int
		expected int
	}{
		{[]int{-1, 2, 1, -4}, 1, 2},
		{[]int{0, 0, 0}, 1, 0},
		{[]int{0, 0, 2, 3, 1}, 2, 2},
	}
	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			out := threeSumClosest(c.nums, c.target)
			require.Equal(t, c.expected, out, "expected %d != out %d", c.expected, out)
		})
	}
}
