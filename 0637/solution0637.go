package solution0637

import (
	"dev/Play-Leetcode-in-Go/kit"
	"dev/Play-with-Data-Structures-in-Go/04-Linked-List/queue"
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

type levelNode struct {
	node  *TreeNode
	level int
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	records := make([][]int, 0)
	q := queue.CreateLoopQueue()
	q.Offer(&levelNode{root, 0})
	for !q.IsEmpty() {
		front := q.Poll().(*levelNode)
		if front.level == len(records) {
			records = append(records, make([]int, 0))
		}
		records[front.level] = append(records[front.level], front.node.Val)
		if front.node.Left != nil {
			q.Offer(&levelNode{front.node.Left, front.level + 1})
		}
		if front.node.Right != nil {
			q.Offer(&levelNode{front.node.Right, front.level + 1})
		}
	}
	res := make([]float64, len(records))
	for i, level := range records {
		res[i] = avg(level)
	}
	return res
}

func avg(nums []int) float64 {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return float64(sum) / float64(len(nums))
}
