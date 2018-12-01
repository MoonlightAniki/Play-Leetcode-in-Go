package _404_v2

import "dev/Play-Leetcode-in-Go/kit"

// 404. Sum of Left Leaves
// https://leetcode.com/problems/sum-of-left-leaves/
/*
Find the sum of all left leaves in a given binary tree.

Example:
    3
   / \
  9  20
    /  \
   15   7
There are two left leaves in the binary tree, with values 9 and 15 respectively. Return 24.
 */

type TreeNode = kit.TreeNode

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil {
		return sumOfLeftLeaves(root.Right)
	}

	if root.Left.Left == nil && root.Left.Right == nil {
		return root.Left.Val + sumOfLeftLeaves(root.Right)
	}

	return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}

// Runtime: 4 ms, faster than 10.53% of Go online submissions for Sum of Left Leaves.
