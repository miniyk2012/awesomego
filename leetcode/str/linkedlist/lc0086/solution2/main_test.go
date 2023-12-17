package main

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/miniyk2012/awesomego/leetcode/utils"
)

type param struct {
	input []int
	x     int
}

type ans struct {
	output []int
}

type question86 struct {
	param
	ans
}

func TestPartition(t *testing.T) {
	qs := []question86{
		{param{input: []int{1, 4, 3, 2, 5, 2}, x: 3},
			ans{output: []int{1, 2, 2, 4, 3, 5}}},
		{param{input: []int{2, 1}, x: 2},
			ans{output: []int{1, 2}}},
	}

	for i, q := range qs {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			head := utils.FromSlice2ListNode(q.input)
			result := partition(head, q.x)
			actual := utils.FromListNode2Slice(result)
			if !reflect.DeepEqual(actual, q.output) {
				t.Fatalf("actual %v != expected %v", actual, q.output)
			}
		})
	}
}
