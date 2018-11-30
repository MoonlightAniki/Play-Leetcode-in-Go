package solution0653

import "dev/Play-Leetcode-in-Go/kit"

// 653. Two Sum IV - Input is a BST
// https://leetcode.com/problems/two-sum-iv-input-is-a-bst/
/*
Given a Binary Search Tree and a target number, return true if there exist two elements in the BST such that their sum is equal to the given target.

Example 1:
Input:
    5
   / \
  3   6
 / \   \
2   4   7
Target = 9
Output: True

Example 2:
Input:
    5
   / \
  3   6
 / \   \
2   4   7
Target = 28
Output: False
 */

type TreeNode = kit.TreeNode

func findTarget(root *TreeNode, k int) bool {
	records := make([]int, 0)
	inOrderTraverse(root, &records)
	i, j := 0, len(records)-1
	for i < j {
		sum := records[i] + records[j]
		if sum == k {
			return true
		} else if sum < k {
			i++
		} else {
			j--
		}
	}
	return false
}

func inOrderTraverse(root *TreeNode, records *[]int) {
	if root == nil {
		return
	}
	inOrderTraverse(root.Left, records)
	*records = append(*records, root.Val)
	inOrderTraverse(root.Right, records)
}

// Runtime: 32 ms, faster than 100.00% of Go online submissions for Two Sum IV - Input is a BST.
