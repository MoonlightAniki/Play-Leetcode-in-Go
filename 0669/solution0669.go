package solution0669

import "dev/Play-Leetcode-in-Go/kit"

// 669. Trim a Binary Search Tree
// https://leetcode.com/problems/trim-a-binary-search-tree/
/*
Given a binary search tree and the lowest and highest boundaries as L and R, trim the tree so that all its elements lies in [L, R] (R >= L). You might need to change the root of the tree, so the result should return the new root of the trimmed binary search tree.

Example 1:
Input:
    1
   / \
  0   2

  L = 1
  R = 2

Output:
    1
      \
       2

Example 2:
Input:
    3
   / \
  0   4
   \
    2
   /
  1

  L = 1
  R = 3

Output:
      3
     /
   2
  /
 1
 */

type TreeNode = kit.TreeNode

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < L {
		return trimBST(root.Right, L, R)
	} else if root.Val > R {
		return trimBST(root.Left, L, R)
	} else {
		root.Left = trimBST(root.Left, L, R)
		root.Right = trimBST(root.Right, L, R)
		return root
	}
}
// Runtime: 16 ms, faster than 100.00% of Go online submissions for Trim a Binary Search Tree.
