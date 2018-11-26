package solution0872_v3

import (
	"dev/Play-Leetcode-in-Go/kit"
	"reflect"
)

// 872. Leaf-Similar Trees
// https://leetcode.com/problems/leaf-similar-trees/
/*
Consider all the leaves of a binary tree.  From left to right order, the values of those leaves form a leaf value sequence.

For example, in the given tree above, the leaf value sequence is (6, 7, 4, 9, 8).

Two binary trees are considered leaf-similar if their leaf value sequence is the same.

Return true if and only if the two given trees with head nodes root1 and root2 are leaf-similar.

Note:

Both of the given trees will have between 1 and 100 nodes.
 */

type TreeNode = kit.TreeNode

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	return reflect.DeepEqual(leaves(root1), leaves(root2))
}

func leaves(root *TreeNode) []int {
	res := make([]int, 0)
	if root.Left != nil {
		res = append(res, leaves(root.Left)...)
	}
	if root.Right != nil {
		res = append(res, leaves(root.Right)...)
	}
	if root.Left == nil && root.Right == nil {
		res = append(res, root.Val)
	}
	return res
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Leaf-Similar Trees.
