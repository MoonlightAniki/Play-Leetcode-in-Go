package solution0001

import "fmt"

// 1. Two Sum
// https://leetcode.com/problems/two-sum/
/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:
Given nums = [2, 7, 11, 15], target = 9,
Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
 */

func twoSum(nums []int, target int) []int {
	records := make(map[int]int)
	for i, v := range nums {
		if j, ok := records[target-v]; ok {
			return []int{j, i}
		} else {
			records[v] = i
		}
	}
	return []int{-1, -1}
}

func Test() {
	nums := []int{2, 7, 11, 15}
	target := 9
	res := twoSum(nums, target)
	fmt.Println(res)
}

// Runtime: 4 ms, faster than 100.00% of Go online submissions for Two Sum.
