package _538

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

func convertBST(root *TreeNode) *TreeNode {
	nums := make([]int, 0)
	inOrderGetVal(root, &nums)
	for i := len(nums) - 2; i >= 0; i-- {
		nums[i] += nums[i+1]
	}
	inOrderSetVal(root, &nums)
	return root
}

func inOrderSetVal(root *TreeNode, nodes *[]int) {
	if root == nil {
		return
	}
	inOrderSetVal(root.Left, nodes)
	root.Val = (*nodes)[0]
	*nodes = (*nodes)[1:]
	inOrderSetVal(root.Right, nodes)
}

func inOrderGetVal(root *TreeNode, nodes *[]int) {
	if root == nil {
		return
	}
	inOrderGetVal(root.Left, nodes)
	*nodes = append(*nodes, root.Val)
	inOrderGetVal(root.Right, nodes)
}
// Runtime: 292 ms, faster than 46.34% of Go online submissions for Convert BST to Greater Tree.
