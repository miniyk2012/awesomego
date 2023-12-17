package main

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/miniyk2012/awesomego/leetcode/utils"
)

type param struct {
	l1 []int
	l2 []int
}

type ans struct {
	result []int
}

type question02 struct {
	param
	ans
}

func TestAddTwoNumbers(t *testing.T) {
	qs := []question02{
		{
			param{[]int{2, 4, 3}, []int{5, 6, 4}},
			ans{[]int{7, 0, 8}},
		},
		{
			param{[]int{0}, []int{0}},
			ans{[]int{0}},
		},
		{
			param{[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}},
			ans{[]int{8, 9, 9, 9, 0, 0, 0, 1}},
		},
	}
	for i, q := range qs {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			l1 := utils.FromSlice2ListNode(q.l1)
			l2 := utils.FromSlice2ListNode(q.l2)
			result := addTwoNumbers(l1, l2)
			actual := utils.FromListNode2Slice(result)
			if !reflect.DeepEqual(actual, q.result) {
				t.Fatalf("expected %v, got %v", q.ans, actual)
			}
		})

	}
}
