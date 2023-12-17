package main

import (
	"github.com/miniyk2012/awesomego/leetcode/utils"
)

// ListNode _
type ListNode = utils.ListNode

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	if k == 0 {
		return head
	}
	var n = 1
	tail := head
	for ; tail.Next != nil; tail = tail.Next {
		n++
	}
	tail.Next = head
	k = k % n
	var before = head
	for i := 0; i < n-k; i++ {
		before = head
		head = head.Next
	}
	before.Next = nil
	return head
}
