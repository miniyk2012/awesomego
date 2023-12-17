package main

import (
	"reflect"
	"strconv"
	"testing"
)

func TestThreeSum(t *testing.T) {
	cases := []struct {
		nums     []int
		expected [][]int
	}{
		{
			[]int{-1, 0, 1, 2, -1, -4},
			[][]int{{-1, 0, 1}, {-1, -1, 2}},
		},
		{
			[]int{0, 1, 1},
			[][]int{},
		},
		{
			[]int{0, 0, 0},
			[][]int{{0, 0, 0}},
		},
		{
			[]int{0, 0, 0, 0},
			[][]int{{0, 0, 0}},
		},
		{
			[]int{-1, 0, 1, 2, -1, -1, -4},
			[][]int{{-1, 0, 1}, {-1, -1, 2}},
		},
	}
	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if actual := threeSum(c.nums); !reflect.DeepEqual(actual, c.expected) {
				t.Fatalf("expected %v, got %v\n", c.expected, actual)
			}
		})
	}
}
