// 参考栈的方法, 并且使用了哑元
// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/solution/shan-chu-lian-biao-de-dao-shu-di-nge-jie-dian-b-61/
package main

import "github.com/miniyk2012/awesomego/leetcode/utils"

// ListNode 对外保持一致
type ListNode = utils.ListNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	var nodes []*ListNode
	for cur := dummy; cur != nil; cur = cur.Next {
		nodes = append(nodes, cur)
	}
	prev := nodes[len(nodes)-1-n]
	prev.Next = prev.Next.Next
	return dummy.Next
}
