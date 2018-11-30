package solution0226

import "dev/Play-Leetcode-in-Go/kit"

// 226. Invert Binary Tree
// https://leetcode.com/problems/invert-binary-tree/
/*
Invert a binary tree.

Example:
Input:
     4
   /   \
  2     7
 / \   / \
1   3 6   9
Output:
     4
   /   \
  7     2
 / \   / \
9   6 3   1

Trivia:
This problem was inspired by this original tweet by Max Howell:
Google: 90% of our engineers use the software you wrote (Homebrew), but you can’t invert a binary tree on a whiteboard so f*** off.
 */

type TreeNode = kit.TreeNode

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := root.Left
	right := root.Right
	root.Left = invertTree(right)
	root.Right = invertTree(left)
	return root
}

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Invert Binary Tree.
