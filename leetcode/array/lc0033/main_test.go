package main

import (
	"strconv"
	"testing"
)

func TestSearch(t *testing.T) {
	cases := []struct {
		nums     []int
		target   int
		expected int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{1}, 0, -1},
		{[]int{3, 5, 1}, 3, 0},
		{[]int{5, 1, 3}, 3, 2},
	}
	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := search(c.nums, c.target)
			if actual != c.expected {
				t.Errorf("expected %d, but got %d", c.expected, actual)
			}
		})

	}
}
