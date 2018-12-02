package _083

import "dev/Play-with-Data-Structures-in-Go/05-Recursion/leetcode/Kit"

// 83. Remove Duplicates from Sorted List
// https://leetcode.com/problems/remove-duplicates-from-sorted-list/
/*
Given a sorted linked list, delete all duplicates such that each element appear only once.

Example 1:
Input: 1->1->2
Output: 1->2

Example 2:
Input: 1->1->2->3->3
Output: 1->2->3
 */
type ListNode = Kit.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prev := head
	for prev.Next != nil {
		if prev.Next.Val == prev.Val {
			delNode := prev.Next
			prev.Next = delNode.Next
			delNode.Next = nil
		} else {
			prev = prev.Next
		}
	}
	return head
}

// Runtime: 4 ms, faster than 100.00% of Go online submissions for Remove Duplicates from Sorted List.
