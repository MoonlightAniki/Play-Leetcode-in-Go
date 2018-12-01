package _543

import "dev/Play-Leetcode-in-Go/kit"

// 543. Diameter of Binary Tree
// https://leetcode.com/problems/diameter-of-binary-tree/
/*
Given a binary tree, you need to compute the length of the diameter of the tree. The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.

Example:
Given a binary tree
          1
         / \
        2   3
       / \
      4   5
Return 3, which is the length of the path [4,2,1,3] or [5,2,1,3].

Note: The length of path between two nodes is represented by the number of edges between them.
 */

type TreeNode = kit.TreeNode

func diameterOfBinaryTree(root *TreeNode) int {
	res := 0
	maxDepth(root, &res)
	return res
}

func maxDepth(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	leftMaxDepth := maxDepth(root.Left, res)
	rightMaxDepth := maxDepth(root.Right, res)
	*res = max(*res, leftMaxDepth+rightMaxDepth)
	return 1 + max(leftMaxDepth, rightMaxDepth)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Runtime: 8 ms, faster than 100.00% of Go online submissions for Diameter of Binary Tree.
