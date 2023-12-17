package main

import (
	"github.com/miniyk2012/awesomego/leetcode/utils"
)

func detectCycle(head *utils.ListNode) *utils.ListNode {
	seen := make(map[*utils.ListNode]struct{})
	current := head
	for current != nil {
		if _, ok := seen[current]; ok {
			return current
		} else {
			seen[current] = struct{}{}
			current = current.Next
		}
	}
	return nil
}
