package solution0872

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
	a := make([]int, 0)
	b := make([]int, 0)
	dfs(root1, &a)
	dfs(root2, &b)
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

func dfs(root *TreeNode, leaves *[]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*leaves = append(*leaves, root.Val)
		return
	}
	dfs(root.Left, leaves)
	dfs(root.Right, leaves)
}

// Runtime: 4 ms, faster than 17.54% of Go online submissions for Leaf-Similar Trees.
