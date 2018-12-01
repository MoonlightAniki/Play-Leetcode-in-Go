package _021

import "dev/Play-with-Data-Structures-in-Go/05-Recursion/leetcode/Kit"

// 21. Merge Two Sorted Lists
// https://leetcode.com/problems/merge-two-sorted-lists/
/*
Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

Example:
Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
 */

type ListNode = Kit.ListNode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	dummyHead := &ListNode{}
	prev := dummyHead

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			node := l1
			l1 = l1.Next
			node.Next = nil
			prev.Next = node
			prev = prev.Next
		} else {
			node := l2
			l2 = l2.Next
			node.Next = nil
			prev.Next = node
			prev = prev.Next
		}
	}
	if l1 != nil {
		prev.Next = l1
	}
	if l2 != nil {
		prev.Next = l2
	}
	return dummyHead.Next
}

// Runtime: 4 ms, faster than 100.00% of Go online submissions for Merge Two Sorted Lists.
