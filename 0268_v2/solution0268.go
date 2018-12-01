package _268_v2

// 268. Missing Number
// https://leetcode.com/problems/missing-number/
/*
Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing from the array.

Example 1:
Input: [3,0,1]
Output: 2

Example 2:
Input: [9,6,4,2,3,5,7,0,1]
Output: 8

Note:
Your algorithm should run in linear runtime complexity. Could you implement it using only constant extra space complexity?
 */
func missingNumber(nums []int) int {
	res := 0 ^ len(nums)
	for i, num := range nums {
		res = res ^ i ^ num
	}
	return res
}

// Runtime: 28 ms, faster than 35.56% of Go online submissions for Missing Number.
