package main

import "github.com/miniyk2012/awesomego/good_example/2-go-callvis/mypkg"

func main() {
	funcs()
	var c calls
	c.execution()
	c.invocation()
}

func funcs() {
	mypkg.Exported()
}

type calls struct{}

func (calls) execution() {
	mypkg.Regular()
}

func (calls) invocation() {
	mypkg.T.Static()
	var i mypkg.Iface = mypkg.T
	i.Dynamic()
}
