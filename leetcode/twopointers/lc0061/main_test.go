package main

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/miniyk2012/awesomego/leetcode/utils"
)

type Param struct {
	head []int
	k    int
}

type Ans struct {
	ans []int
}

type questions61 struct {
	Param
	Ans
}

func TestRotateRight(t *testing.T) {
	qs := []questions61{
		{
			Param{head: []int{1, 2, 3, 4, 5}, k: 2},
			Ans{[]int{4, 5, 1, 2, 3}},
		},
		{
			Param{head: []int{0, 1, 2}, k: 0},
			Ans{[]int{0, 1, 2}},
		},
		{
			Param{head: []int{0, 1, 2}, k: 1},
			Ans{[]int{2, 0, 1}},
		},
		{
			Param{head: []int{0, 1, 2}, k: 2},
			Ans{[]int{1, 2, 0}},
		},
		{
			Param{head: []int{0, 1, 2}, k: 3},
			Ans{[]int{0, 1, 2}},
		},
		{
			Param{head: []int{0, 1, 2}, k: 4},
			Ans{[]int{2, 0, 1}},
		},
	}
	for i, q := range qs {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			l := utils.FromSlice2ListNode(q.head)
			res := rotateRight(l, q.k)
			actual := utils.FromListNode2Slice(res)
			if !reflect.DeepEqual(actual, q.ans) {
				t.Errorf("expected %v, got %v", q.ans, actual)
			}
		})
	}
}
