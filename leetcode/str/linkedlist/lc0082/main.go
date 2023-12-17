package main

import (
	"github.com/miniyk2012/awesomego/leetcode/utils"
)

// ListNode _
type ListNode = utils.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Val: 0, Next: head}
	prev, cur := dummy, dummy.Next
	onlyOne := true
	for ; cur.Next != nil; cur = cur.Next {
		if cur.Val != cur.Next.Val {
			if onlyOne {
				prev.Next, prev = cur, cur
			}
			onlyOne = true
		} else {
			onlyOne = false
		}
	}
	if onlyOne {
		prev.Next = cur
	} else {
		prev.Next = nil
	}
	return dummy.Next
}
