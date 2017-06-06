package leetcode

import (
	"testing"
)

//Given a linked list, remove the nth node from the end of list and return its head.
//
//For example,
//
//Given linked list: 1->2->3->4->5, and n = 2.
//
//After removing the second node from the end, the linked list becomes 1->2->3->5.
//Note:
//Given n will always be valid.
//Try to do this in one pass.
//https://leetcode.com/problems/remove-nth-node-from-end-of-list/#/description
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	temp := head
	list := []*ListNode{}
	for temp != nil {
		list = append(list, temp)
		temp = temp.Next
	}
	length := len(list)
	thisOne := list[length - n]
	beforeIdx := length - n - 1;
	if beforeIdx < 0 {
		return thisOne.Next
	}
	list[beforeIdx].Next = thisOne.Next
	return head
}

func Test19(t *testing.T) {

	var lists = []struct {
		head   *ListNode
		n      int
		result string
	}{
		{
			&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			2,
			"[ 1 2 3 5 ]",
		},
	}
	t.Log("Begin Tests")
	for i, l := range lists {
		t.Log("begin test #", i)
		t.Log("removeNthFromEnd", l.head.String(), "hope output:", l.result)
		result := removeNthFromEnd(l.head, l.n)
		t.Log("removeNthFromEnd result:", result)
		if l.result != result.String() {
			t.Fatal("result error, want", l.result, "but output", result)
		}
		t.Log("end test #", i)
	}
	t.Log("End Tests")

}
