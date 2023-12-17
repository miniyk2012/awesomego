package utils

import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	val := "["
	current := l
	for current != nil {
		val += fmt.Sprintf("%d,", current.Val)
		current = current.Next
	}
	val += "]"
	return val
}

// FromSlice2ListNode 由slice构建ListNode
func FromSlice2ListNode(s []int) *ListNode {
	var header, tail *ListNode = nil, nil
	for i := range s {
		if header == nil {
			header = &ListNode{Val: s[i]}
			tail = header
			continue
		}
		tail.Next = &ListNode{Val: s[i]}
		tail = tail.Next
	}
	return header
}

// FromListNode2Slice 由ListNode构建slice
func FromListNode2Slice(l *ListNode) []int {
	s := make([]int, 0)
	for l != nil {
		s = append(s, l.Val)
		l = l.Next
	}
	return s
}
