package main

import (
	"reflect"
	"strconv"
	"testing"
)

func TestClosestPrimes(t *testing.T) {
	testcases := []struct {
		left, right int
		ans         []int
	}{
		{10, 19, []int{11, 13}},
		{4, 6, []int{-1, -1}},
		{19, 31, []int{29, 31}},
	}
	for i, tc := range testcases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ret := closestPrimes(tc.left, tc.right)
			if !reflect.DeepEqual(ret, tc.ans) {
				t.Fatalf("expected %v, got %v", tc.ans, ret)
			}
		})
	}
}
