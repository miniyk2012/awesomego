package main

import "github.com/miniyk2012/awesomego/leetcode/utils"

// ListNode is singly-linked list
type ListNode = utils.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return head
	}
	dump := &ListNode{0, head}
	before := dump
	i := 0
	for ; i < left-1; i++ {
		before = before.Next
	}
	pre, next := before.Next, before.Next.Next
	for ; i < right-1; i++ {
		next.Next, pre, next = pre, next, next.Next
	}
	before.Next, before.Next.Next = pre, next
	return dump.Next
}
