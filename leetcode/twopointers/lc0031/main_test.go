package main

import (
	"reflect"
	"strconv"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	cases := []struct {
		nums     []int
		expected []int
	}{
		{[]int{1, 3, 2}, []int{2, 1, 3}},
		{[]int{1, 2, 3}, []int{1, 3, 2}},
		{[]int{1, 1, 5}, []int{1, 5, 1}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
	}
	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			nextPermutation(c.nums)
			if !reflect.DeepEqual(c.nums, c.expected) {
				t.Fatalf("expected %v, got %v", c.expected, c.nums)
			}
		})

	}

}
