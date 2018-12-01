package _404

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

var sum = 0

func sumOfLeftLeaves(root *TreeNode) int {
	sum = 0
	dfs(root)
	return sum
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	dfs(root.Right)

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		sum += root.Left.Val
	}
}

// Runtime: 4 ms, faster than 10.53% of Go online submissions for Sum of Left Leaves.
