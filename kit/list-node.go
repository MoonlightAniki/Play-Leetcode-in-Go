package kit

import (
	"bytes"
	"fmt"
)

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	dummy := &ListNode{}
	prev := dummy
	for _, num := range nums {
		prev.Next = &ListNode{num, nil}
		prev = prev.Next
	}
	return dummy.Next
}

func (head *ListNode) String() string {
	var buffer bytes.Buffer
	for cur := head; cur != nil; cur = cur.Next {
		buffer.WriteString(fmt.Sprintf("%d->", cur.Val))
	}
	buffer.WriteString("nil")
	return buffer.String()
}
