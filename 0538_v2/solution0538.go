package _538_v2

import "dev/Play-Leetcode-in-Go/kit"

// 538. Convert BST to Greater Tree
// https://leetcode.com/problems/convert-bst-to-greater-tree/
/*
Given a Binary Search Tree (BST), convert it to a Greater Tree such that every key of the original BST is changed to the
original key plus sum of all keys greater than the original key in BST.

Example:
Input: The root of a Binary Search Tree like this:
              5
            /   \
           2     13

Output: The root of a Greater Tree like this:
             18
            /   \
          20     13
 */

type TreeNode = kit.TreeNode

var sum = 0

func convertBST(root *TreeNode) *TreeNode {
	sum = 0
	dfs(root)
	return root
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Right)
	sum += root.Val
	root.Val = sum
	dfs(root.Left)
}

// Runtime: 272 ms, faster than 73.17% of Go online submissions for Convert BST to Greater Tree.
