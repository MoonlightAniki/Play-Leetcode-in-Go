package solution0876

import (
	"dev/Play-Leetcode-in-Go/kit"
	"fmt"
)

// 700. Search in a Binary Search Tree
// https://leetcode.com/problems/middle-of-the-linked-list/
/*
Given a non-empty, singly linked list with head node head, return a middle node of linked list.
If there are two middle nodes, return the second middle node.

Example 1:
Input: [1,2,3,4,5]
Output: Node 3 from this list (Serialization: [3,4,5])
The returned node has value 3.  (The judge's serialization of this node is [3,4,5]).
Note that we returned a ListNode object ans, such that:
ans.val = 3, ans.next.val = 4, ans.next.next.val = 5, and ans.next.next.next = NULL.

Example 2:
Input: [1,2,3,4,5,6]
Output: Node 4 from this list (Serialization: [4,5,6])
Since the list has two middle nodes with values 3 and 4, we return the second one.

Note:
The number of nodes in the given list will be between 1 and 100.
 */

type ListNode = kit.ListNode

func middleNode(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func Test() {
	head := kit.CreateListNode([]int{1, 2, 3})
	fmt.Println(head)
	fmt.Println(middleNode(head))
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Middle of the Linked List.
