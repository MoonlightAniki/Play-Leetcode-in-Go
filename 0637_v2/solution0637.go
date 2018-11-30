package solution0637_v2

import (
	"dev/Play-Leetcode-in-Go/kit"
)

// 637. Average of Levels in Binary Tree
// https://leetcode.com/problems/average-of-levels-in-binary-tree/
/*
Given a non-empty binary tree, return the average value of the nodes on each level in the form of an array.
Example 1:
Input:
    3
   / \
  9  20
    /  \
   15   7
Output: [3, 14.5, 11]
Explanation:
The average value of nodes on level 0 is 3,  on level 1 is 14.5, and on level 2 is 11. Hence return [3, 14.5, 11].
Note:
The range of node's value is in the range of 32-bit signed integer.
 */

type TreeNode = kit.TreeNode

func averageOfLevels(root *TreeNode) []float64 {
	res := make([]float64, 0, 128)

	nodes := make([]*TreeNode, 0, 1024)
	nodes = append(nodes, root)

	for len(nodes) > 0 {
		n := len(nodes)

		sum := 0
		for i := 0; i < n; i++ {
			sum += nodes[i].Val
			if nodes[i].Left != nil {
				nodes = append(nodes, nodes[i].Left)
			}
			if nodes[i].Right != nil {
				nodes = append(nodes, nodes[i].Right)
			}
		}
		res = append(res, float64(sum)/float64(n))
		nodes = nodes[n:]
	}
	return res
}
// Runtime: 16 ms, faster than 100.00% of Go online submissions for Average of Levels in Binary Tree.
