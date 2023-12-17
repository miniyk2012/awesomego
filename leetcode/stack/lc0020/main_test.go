package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question20 struct {
	para20
	ans20
}

// para 是参数
// one 代表第一个参数
type para20 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans20 struct {
	one bool
}

func TestIsValid(t *testing.T) {
	qs := []question20{
		{
			para20{"()[]{}"},
			ans20{true},
		},
		{
			para20{"(]"},
			ans20{false},
		},
		{
			para20{"({[]})"},
			ans20{true},
		},
		{
			para20{"(){[({[]})]}"},
			ans20{true},
		},
		{
			para20{"((([[[{{{"},
			ans20{false},
		},
		{
			para20{"(())]]"},
			ans20{false},
		},
		{
			para20{""},
			ans20{true},
		},
		{
			para20{"["},
			ans20{false},
		},
	}
	for i, q := range qs {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			a, p := q.ans20, q.para20
			assert.Equal(t, a.one, isValid(p.one))
		})
	}
}
