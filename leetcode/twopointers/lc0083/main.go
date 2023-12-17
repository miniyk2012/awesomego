package main

import "github.com/miniyk2012/awesomego/leetcode/utils"

// ListNode _
type ListNode = utils.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur, before := head.Next, head
	for ; cur.Next != nil; cur = cur.Next {
		if before.Val == cur.Val {
			before.Next = cur.Next
		} else {
			before = cur
		}
	}
	if before.Val == cur.Val {
		before.Next = nil
	}
	return head
}
