package main

import "github.com/miniyk2012/awesomego/leetcode/utils"

// ListNode _
type ListNode = utils.ListNode

func partition(head *ListNode, x int) *ListNode {
	left := &ListNode{Val: 0, Next: nil}
	leftCur := left
	right := &ListNode{Val: 0, Next: nil}
	rightCur := right
	for cur := head; cur != nil; cur = cur.Next {
		if cur.Val < x {
			leftCur.Next = cur
			leftCur = cur
		} else {
			rightCur.Next = cur
			rightCur = cur
		}
	}
	rightCur.Next = nil
	leftCur.Next = right.Next
	return left.Next
}
