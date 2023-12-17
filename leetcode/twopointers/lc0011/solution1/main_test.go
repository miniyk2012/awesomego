package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArea(t *testing.T) {
	cases := []struct {
		height   []int
		expected int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
	}
	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			out := maxArea(c.height)
			require.Equal(t, c.expected, out, "expected %d != out %d", c.expected, out)
		})
	}
}
