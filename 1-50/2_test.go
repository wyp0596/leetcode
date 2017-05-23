package leetcode

import (
	"testing"
)

//You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
//
//You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//
//Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
//Output: 7 -> 0 -> 8
//https://leetcode.com/problems/add-two-numbers/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	result := &ListNode{-1, nil}
	curNode := result
	overflow := false
	for l1 != nil || l2 != nil || overflow {
		sum := 0
		if overflow {
			sum++
			overflow = false
		}
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		if sum >= 10 {
			sum -= 10
			overflow = true
		}
		if curNode.Val >= 0 {
			curNode.Next = new(ListNode)
			curNode = curNode.Next
		}
		curNode.Val = sum
	}
	return result
}

func Test2(t *testing.T) {

	var lists = []struct {
		l1     *ListNode
		l2     *ListNode
		result string
	}{
		{
			nil,
			nil,
			"[ ]",
		},
		{
			&ListNode{1, &ListNode{3, &ListNode{5, nil}}},
			nil,
			"[ 1 3 5 ]",
		},
		{
			nil,
			&ListNode{2, &ListNode{4, nil}},
			"[ 2 4 ]",
		},
		{
			&ListNode{1, &ListNode{3, &ListNode{5, nil}}},
			&ListNode{2, &ListNode{4, nil}},
			"[ 3 7 5 ]",
		},
		{
			&ListNode{2, &ListNode{6, &ListNode{5, nil}}},
			&ListNode{2, &ListNode{6, nil}},
			"[ 4 2 6 ]",
		},
	}
	t.Log("Begin Tests")
	for i, l := range lists {
		t.Log("begin test #", i)
		t.Log("addTwoNumbers", l.l1, "and", l.l2, "hope output:", l.result)
		result := addTwoNumbers(l.l1, l.l2)
		t.Log("addTwoNumbers result:", result)
		if l.result != result.String() {
			t.Fatal("result error, want", l.result, "but output", result)
		}
		t.Log("end test #", i)
	}
	t.Log("End Tests")

}
