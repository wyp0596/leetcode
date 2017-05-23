package leetcode

import "strconv"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode)String() string {
	s := l
	r := "["
	for s != nil {
		r += " " + strconv.Itoa(s.Val)
		s = s.Next
	}
	r += " ]"
	return r
}
