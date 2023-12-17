package main

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	cases := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{"1",
			[]int{-1, 0, 1, 2, -1, -4},
			[][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{"2",
			[]int{0, 1, 1},
			[][]int{},
		},
		{"3",
			[]int{0, 0, 0},
			[][]int{{0, 0, 0}},
		},
		{"4",
			[]int{0, 0, 0, 0},
			[][]int{{0, 0, 0}},
		},
		{"5",
			[]int{-1, 0, 1, 2, -1, -1, -4},
			[][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if actual := threeSum(c.nums); !reflect.DeepEqual(actual, c.expected) {
				t.Fatalf("expected %v, got %v\n", c.expected, actual)
			}
		})
	}
}
