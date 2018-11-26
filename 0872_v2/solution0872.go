package solution0872_v2

import (
	"dev/Play-Leetcode-in-Go/kit"
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
	a := dfs(root1, make([]int, 0))
	b := dfs(root2, make([]int, 0))
	return equals(a, b)
}

func equals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func dfs(root *TreeNode, leaves []int) []int {
	if root == nil {
		return leaves
	}
	if root.Left == nil && root.Right == nil {
		leaves = append(leaves, root.Val)
		return leaves
	}
	leaves = dfs(root.Left, leaves)
	leaves = dfs(root.Right, leaves)
	return leaves
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Leaf-Similar Trees.
