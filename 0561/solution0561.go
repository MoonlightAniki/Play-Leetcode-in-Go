package solution561

// 561. Array Partition I
// https://leetcode.com/problems/array-partition-i/
/*
Given an array of 2n integers, your task is to group these integers into n pairs of integer, say (a1, b1), (a2, b2), ..., (an, bn)
which makes sum of min(ai, bi) for all i from 1 to n as large as possible.

Example 1:
Input: [1,4,3,2]
Output: 4
Explanation: n is 2, and the maximum sum of pairs is 4 = min(1, 2) + min(3, 4).

Note:
n is a positive integer, which is in the range of [1, 10000].
All the integers in the array will be in the range of [-10000, 10000].
 */

func arrayPairSum(nums []int) int {
	shellSort(nums)
	res := 0
	for i := 0; i < len(nums); i += 2 {
		res += nums[i]
	}
	return res
}

// 希尔排序
func shellSort(nums []int) {
	h := 1
	for h < len(nums)/3 {
		h = 3*h + 1
	}
	for h > 0 {
		for i := h; i < len(nums); i += h {
			v := nums[i]
			j := i
			for ; j-h >= 0 && nums[j-h] > v; j -= h {
				nums[j] = nums[j-h]
			}
			nums[j] = v
		}
		h /= 3
	}
}

// Runtime: 424 ms, faster than 2.06% of Go online submissions for Array Partition I.
