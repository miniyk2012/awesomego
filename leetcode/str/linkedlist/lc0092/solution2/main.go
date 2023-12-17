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

// 头插法
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return head
	}
	dump := &ListNode{0, head}
	pre := dump
	for i := 0; i < left-1; i++ {
		pre = pre.Next // pre是左边节点的前一个节点, 永远不变了
	}
	cur := pre.Next // cur指向已插入的节点末尾
	for i := 0; i < right-left; i++ {
		next := cur.Next // next此时表示要插入的节点(准备放到头部去)
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dump.Next
}
