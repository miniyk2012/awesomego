package main

import "github.com/miniyk2012/awesomego/leetcode/utils"

func removeNthFromEnd(head *utils.ListNode, n int) *utils.ListNode {
	right, left := head, head
	for i := 0; i < n; i++ {
		right = right.Next
	}
	if right == nil {
		head.Next, head = nil, head.Next
		return head
	}
	for right.Next != nil {
		right = right.Next
		left = left.Next
	}
	removeNode := left.Next
	removeNode.Next, left.Next = nil, removeNode.Next
	return head
}
