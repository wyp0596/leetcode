package leetcode_test

import (
	"testing"
	"reflect"
)

func TestMerge(t *testing.T) {

	var lists = []struct {
		l1     *ListNode
		l2     *ListNode
		result []int
	}{
		{
			nil,
			nil,
			nil,
		},
		{
			&ListNode{1, &ListNode{3, &ListNode{5, nil}}},
			nil,
			[]int{1, 3, 5},
		},
		{
			nil,
			&ListNode{2, &ListNode{4, nil}},
			[]int{2, 4},
		},
		{
			&ListNode{1, &ListNode{3, &ListNode{5, nil}}},
			&ListNode{2, &ListNode{4, nil}},
			[]int{1, 2, 3, 4, 5},
		},
		{
			&ListNode{2, &ListNode{3, &ListNode{5, nil}}},
			&ListNode{2, &ListNode{6, nil}},
			[]int{2, 2, 3, 5, 6},
		},
	}
	t.Log("Begin Tests")
	for i, l := range lists {
		t.Log("begin test", i)
		t.Log("merge", l.l1, "and", l.l2)
		result := mergeTwoLists(l.l1, l.l2)
		t.Log("merge result:", result)
		r := getListNode(result)
		if !reflect.DeepEqual(r, l.result) {
			t.Fatal("result error, want", l.result, "but output", r)
		}
		t.Log("end test", i)
	}
	t.Log("End Tests")

}

func getListNode(l *ListNode) []int {
	var r []int
	for l != nil {
		r = append(r, l.Val)
		l = l.Next
	}
	return r
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}
// Merge two sorted linked lists and return it as a new list.
// The new list should be made by splicing together the nodes of the first two lists.
// https://leetcode.com/problems/merge-two-sorted-lists/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var result *ListNode
	if l1.Val < l2.Val {
		result = l1
		l1 = l1.Next
	} else {
		result = l2
		l2 = l2.Next
	}
	current := result
	for l1 != nil {
		if l2 != nil && l1.Val > l2.Val {
			current.Next = l2
			current = l2
			l2 = l2.Next
			continue
		}
		current.Next = l1
		current = l1
		l1 = l1.Next
	}
	if l2 != nil {
		current.Next = l2
	}

	return result
}