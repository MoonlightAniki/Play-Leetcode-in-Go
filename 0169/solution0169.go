package _169

// 169. Majority Element
// https://leetcode.com/problems/majority-element/
/*
Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.
You may assume that the array is non-empty and the majority element always exist in the array.

Example 1:
Input: [3,2,3]
Output: 3

Example 2:
Input: [2,2,1,1,1,2,2]
Output: 2
 */

func majorityElement(nums []int) int {
	n, freq := len(nums), make(map[int]int, len(nums))
	for _, num := range nums {
		freq[num]++
		if freq[num] > n/2 {
			return num
		}
	}
	return 0
}

// Runtime: 28 ms, faster than 77.85% of Go online submissions for Majority Element.
