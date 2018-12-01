package _563

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
	tilt = 0
	m = make(map[*TreeNode]int)
	getTilt(root)
	return tilt
}

var tilt = 0
var m = make(map[*TreeNode]int)

func getTilt(root *TreeNode) {
	if root == nil {
		return
	}
	getTilt(root.Left)
	getTilt(root.Right)
	tilt += abs(getSum(root.Left) - getSum(root.Right))
}

func getSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if sum, ok := m[root]; ok {
		return sum
	}
	m[root] = root.Val + getSum(root.Left) + getSum(root.Right)
	return m[root]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Runtime: 20 ms, faster than 60.00% of Go online submissions for Binary Tree Tilt.
