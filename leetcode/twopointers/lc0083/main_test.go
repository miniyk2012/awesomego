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

type question83 struct {
	param
	ans
}

func TestDeleteDuplicates(t *testing.T) {
	qs := []question83{
		{
			param{[]int{1, 2, 3}},
			ans{[]int{1, 2, 3}},
		},
		{
			param{[]int{1, 2, 3, 3, 4, 4, 5}},
			ans{[]int{1, 2, 3, 4, 5}},
		},
		{
			param{[]int{1, 1, 1, 2, 3}},
			ans{[]int{1, 2, 3}},
		},
		{
			param{[]int{1, 1, 2, 3, 3}},
			ans{[]int{1, 2, 3}},
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
