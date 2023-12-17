package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question28 struct {
	para28
	ans28
}

type para28 struct {
	s1, s2 string
}

type ans28 struct {
	a int
}

func TestStrStr(t *testing.T) {
	qs := []question28{
		{
			para28{
				s1: "sadbutsad", s2: "sad",
			},
			ans28{
				a: 0,
			},
		},
		{
			para28{
				s1: "leetcode", s2: "leeto",
			},
			ans28{
				a: -1,
			},
		},
	}
	for i, q := range qs {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, strStr(q.s1, q.s2), q.a)
		})
	}

}
