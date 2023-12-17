package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type param struct {
	nums []int
}

type ans struct {
	k        int
	expected []int
}

type question83 struct {
	param
	ans
}

func TestDeleteDuplicates(t *testing.T) {
	qs := []question83{
		{
			param{[]int{1, 1, 1, 2, 2, 3}},
			ans{5, []int{1, 1, 2, 2, 3}},
		},
		{
			param{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}},
			ans{7, []int{0, 0, 1, 1, 2, 3, 3}},
		},
		{
			param{[]int{0}},
			ans{1, []int{0}},
		},
	}
	for i, q := range qs {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := removeDuplicates(q.nums)
			assert.Equal(t, q.k, r)
			for i := 0; i < q.k; i++ {
				assert.Equal(t, q.expected[i], q.nums[i])
			}
		})
	}
}
