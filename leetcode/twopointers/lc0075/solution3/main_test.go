package main

import (
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func TestSortColors(t *testing.T) {
	cases := []struct {
		nums []int
	}{
		{[]int{2, 0, 2, 1, 1, 0}},
		{[]int{2, 0, 1}},
	}
	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			sortColors(c.nums)
			expected := make([]int, len(c.nums))
			copy(expected, c.nums)
			sort.Ints(expected)
			if !reflect.DeepEqual(c.nums, expected) {
				t.Errorf("expected: %v, got: %v", expected, c.nums)
			}
		})
	}
}
