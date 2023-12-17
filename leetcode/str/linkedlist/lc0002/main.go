package main

import (
	"github.com/miniyk2012/awesomego/leetcode/utils"
)

// ListNode _
type ListNode = utils.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	cur := dummy
	left := l1
	right := l2
	carry := 0
	for ; left != nil && right != nil; left, right = left.Next, right.Next {
		total := left.Val + right.Val + carry
		val := total % 10
		carry = total / 10
		cur.Next = &ListNode{Val: val}
		cur = cur.Next
	}
	for ; left != nil; left = left.Next {
		total := left.Val + carry
		val := total % 10
		carry = total / 10
		cur.Next = &ListNode{Val: val}
		cur = cur.Next
	}
	for ; right != nil; right = right.Next {
		total := right.Val + carry
		val := total % 10
		carry = total / 10
		cur.Next = &ListNode{Val: val}
		cur = cur.Next
	}
	if carry > 0 {
		cur.Next = &ListNode{Val: carry}
		cur = cur.Next
	}
	return dummy.Next
}
