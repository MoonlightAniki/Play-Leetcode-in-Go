package _572

import "dev/Play-Leetcode-in-Go/kit"

// 572. Subtree of Another Tree
// https://leetcode.com/problems/subtree-of-another-tree/
/*
Given two non-empty binary trees s and t, check whether tree t has exactly the same structure and node values with a
subtree of s. A subtree of s is a tree consists of a node in s and all of this node's descendants. The tree s could
also be considered as a subtree of itself.

Example 1:
Given tree s:
     3
    / \
   4   5
  / \
 1   2
Given tree t:
   4
  / \
 1   2
Return true, because t has the same structure and node values with a subtree of s.

Example 2:
Given tree s:
     3
    / \
   4   5
  / \
 1   2
    /
   0
Given tree t:
   4
  / \
 1   2
Return false.
 */
type TreeNode = kit.TreeNode

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return t == nil
	}
	return isSame(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSame(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	return s.Val == t.Val && isSame(s.Left, t.Left) && isSame(s.Right, t.Right)
}

// Runtime: 24 ms, faster than 100.00% of Go online submissions for Subtree of Another Tree.
