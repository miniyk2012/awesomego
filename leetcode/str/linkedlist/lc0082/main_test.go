package main

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/miniyk2012/awesomego/leetcode/utils"
)

type param struct {
	head []int
}

type ans struct {
	expected []int
}

type question82 struct {
	param
	ans
}

func TestDeleteDuplicates(t *testing.T) {
	qs := []question82{
		{
			param{[]int{1, 2, 3}},
			ans{[]int{1, 2, 3}},
		},
		{
			param{[]int{0}},
			ans{[]int{0}},
		},
		{
			param{[]int{1, 2, 3, 3, 4, 4, 5}},
			ans{[]int{1, 2, 5}},
		},
		{
			param{[]int{1, 1, 1, 2, 3}},
			ans{[]int{2, 3}},
		},
		{
			param{[]int{1, 1, 2, 3, 3}},
			ans{[]int{2}},
		},
		{
			param{[]int{1, 1, 3, 3}},
			ans{[]int{}},
		},
	}
	for i, q := range qs {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			head := utils.FromSlice2ListNode(q.head)
			result := deleteDuplicates(head)
			actual := utils.FromListNode2Slice(result)
			if !reflect.DeepEqual(actual, q.expected) {
				t.Errorf("expected: %v, actual: %v", q.expected, actual)
			}
		})
	}
}
