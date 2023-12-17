package main

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/miniyk2012/awesomego/leetcode/utils"
)

func TestRemoveNthFromEnd(t *testing.T) {
	cases := []struct {
		s        []int
		n        int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
		{[]int{1}, 1, []int{}},
		{[]int{1, 2}, 1, []int{1}},
	}
	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			head := utils.FromSlice2ListNode(c.s)
			l := removeNthFromEnd(head, c.n)
			actual := utils.FromListNode2Slice(l)
			if !reflect.DeepEqual(actual, c.expected) {
				t.Fatalf("expected %v, got %v", c.expected, actual)
			}
		})

	}

}
