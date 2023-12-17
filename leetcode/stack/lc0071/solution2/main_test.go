package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question71 struct {
	para71
	ans71
}

// para 是参数
// s 代表第一个参数
type para71 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans71 struct {
	one string
}

func TestSimplifyPath(t *testing.T) {
	qs := []question71{

		{
			para71{"/.hidden"},
			ans71{"/.hidden"},
		},

		{
			para71{"/..hidden"},
			ans71{"/..hidden"},
		},

		{
			para71{"/abc/..."},
			ans71{"/abc/..."},
		},

		{
			para71{"/home/"},
			ans71{"/home"},
		},

		{
			para71{"/..."},
			ans71{"/..."},
		},

		{
			para71{"/../"},
			ans71{"/"},
		},

		{
			para71{"/home//foo/"},
			ans71{"/home/foo"},
		},

		{
			para71{"/a/./b/../../c/"},
			ans71{"/c"},
		},
	}
	for i, q := range qs {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, q.one, simplifyPath(q.s))
		})
	}

}
