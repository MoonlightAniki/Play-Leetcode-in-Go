package _563_v2

import (
	"dev/Play-Leetcode-in-Go/kit"
)

// 563. Binary Tree Tilt
// https://leetcode.com/problems/binary-tree-tilt/
/*
Given a binary tree, return the tilt of the whole tree.
The tilt of a tree node is defined as the absolute difference between the sum of all left subtree node values and the
sum of all right subtree node values. Null node has tilt 0.
The tilt of the whole tree is defined as the sum of all nodes' tilt.

Example:
Input:
         1
       /   \
      2     3
Output: 1
Explanation:
Tilt of node 2 : 0
Tilt of node 3 : 0
Tilt of node 1 : |2-3| = 1
Tilt of binary tree : 0 + 0 + 1 = 1

Note:
The sum of node values in any subtree won't exceed the range of 32-bit integer.
All the tilt values won't exceed the range of 32-bit integer.
 */

type TreeNode = kit.TreeNode

func findTilt(root *TreeNode) int {
	tilt := 0
	getSum(root, &tilt)
	return tilt
}

func getSum(root *TreeNode, tilt *int) int {
	if root == nil {
		return 0
	}
	l := getSum(root.Left, tilt)
	r := getSum(root.Right, tilt)

	if l > r {
		*tilt += l - r
	} else {
		*tilt += r - l
	}

	return l + r + root.Val
}

// Runtime: 16 ms, faster than 100.00% of Go online submissions for Binary Tree Tilt.
