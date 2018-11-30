package _206

import "dev/Play-with-Data-Structures-in-Go/05-Recursion/leetcode/Kit"

// 206. Reverse Linked List
// https://leetcode.com/problems/reverse-linked-list/
/*
Reverse a singly linked list.

Example:
Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL

Follow up:
A linked list can be reversed either iteratively or recursively. Could you implement both?
 */

type ListNode = Kit.ListNode

func reverseList(head *ListNode) *ListNode {
	dummyHead := &ListNode{}
	for head != nil {
		node := head
		head = head.Next
		node.Next = nil

		node.Next = dummyHead.Next
		dummyHead.Next = node
	}
	return dummyHead.Next
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Linked List.
