package main

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/miniyk2012/awesomego/leetcode/utils"
)

type param struct {
	head        []int
	left, right int
}

type ans struct {
	output []int
}

type question92 struct {
	param
	ans
}

func TestReverseBetween(t *testing.T) {
	qs := []question92{
		{param{head: []int{1, 2, 3, 4, 5}, left: 2, right: 4},
			ans{output: []int{1, 4, 3, 2, 5}}},
		{param{head: []int{5}, left: 1, right: 1},
			ans{output: []int{5}}},
	}
	for i, q := range qs {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			head := utils.FromSlice2ListNode(q.head)
			result := reverseBetween(head, q.left, q.right)
			actual := utils.FromListNode2Slice(result)
			if !reflect.DeepEqual(actual, q.output) {
				t.Fatalf("actual %v != expected %v", actual, q.output)
			}
		})
	}
}
