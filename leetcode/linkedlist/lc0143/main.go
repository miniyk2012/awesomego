package main

import (
	"github.com/miniyk2012/awesomego/leetcode/utils"
)

type ListNode = utils.ListNode

func reorderList(head *ListNode) {
	slice := make([]*ListNode, 0)
	for current := head; current != nil; current = current.Next {
		slice = append(slice, current)
	}
	length := len(slice)
	i, j := 0, length-1
	for i < j {
		slice[i].Next = slice[j]
		i++
		if i == j {
			break
		}
		slice[j].Next = slice[i]
		j--
	}
	slice[i].Next = nil
}
