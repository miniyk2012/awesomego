package main

import "github.com/miniyk2012/awesomego/leetcode/utils"

// ListNode _
type ListNode = utils.ListNode

func partition(head *ListNode, x int) *ListNode {
	left := make([]*ListNode, 0)
	right := make([]*ListNode, 0)

	for cur := head; cur != nil; cur = cur.Next {
		if cur.Val < x {
			left = append(left, cur)
		} else {
			right = append(right, cur)
		}
	}
	dummy := &ListNode{Val: 0, Next: nil}
	cur := dummy
	for _, node := range left {
		cur.Next = node
		cur = node
	}
	for _, node := range right {
		cur.Next = node
		cur = node
	}
	cur.Next = nil
	return dummy.Next
}
