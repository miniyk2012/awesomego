package main

import (
	"testing"

	"github.com/miniyk2012/awesomego/leetcode/utils"
	"github.com/stretchr/testify/assert"
)

func TestReorderList(t *testing.T) {
	testcases := []struct {
		input  []int
		output []int
	}{
		{input: []int{1, 2, 3, 4}, output: []int{1, 4, 2, 3}},
		{input: []int{1, 2, 3, 4, 5}, output: []int{1, 5, 2, 4, 3}},
	}
	for _, testcase := range testcases {
		inputList := utils.FromSlice2ListNode(testcase.input)
		reorderList(inputList)
		recorder := utils.FromListNode2Slice(inputList)
		assert.Equal(t, recorder, testcase.output)
	}
}
