package _437

import "dev/Play-Leetcode-in-Go/kit"

// 437. Path Sum III
// https://leetcode.com/problems/path-sum-iii/
/*
You are given a binary tree in which each node contains an integer value.

Find the number of paths that sum to a given value.

The path does not need to start or end at the root or a leaf, but it must go downwards (traveling only from parent nodes to child nodes).

The tree has no more than 1,000 nodes and the values are in the range -1,000,000 to 1,000,000.

Example:

root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8

      10
     /  \
    5   -3
   / \    \
  3   2   11
 / \   \
3  -2   1

Return 3. The paths that sum to 8 are:

1.  5 -> 3
2.  5 -> 2 -> 1
3. -3 -> 11
 */

type TreeNode = kit.TreeNode

func pathSum(root *TreeNode, sum int) int {
	res := 0
	inOrder(root, sum, &res)
	return res
}

func inOrder(root *TreeNode, sum int, res *int) {
	if root == nil {
		return
	}
	inOrder(root.Left, sum, res)
	__pathSum(root, sum, 0, res)
	inOrder(root.Right, sum, res)
}

func __pathSum(root *TreeNode, sum int, curSum int, res *int) {
	if root == nil {
		return
	}
	curSum += root.Val
	if sum == curSum {
		*res++
	}
	__pathSum(root.Left, sum, curSum, res)
	__pathSum(root.Right, sum, curSum, res)
}

// Runtime: 20 ms, faster than 76.47% of Go online submissions for Path Sum III.
