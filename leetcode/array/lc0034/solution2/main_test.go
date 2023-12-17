package main

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSearchRange(t *testing.T) {
	cases := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{[]int{}, 0, []int{-1, -1}},
	}
	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := searchRange(c.nums, c.target)
			if !reflect.DeepEqual(result, c.expected) {
				t.Fatalf("expected %v, got %v", c.expected, result)
			}
		})
	}
}
