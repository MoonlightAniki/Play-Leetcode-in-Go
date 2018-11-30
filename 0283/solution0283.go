package solution0283

// 283. Move Zeroes
// https://leetcode.com/problems/move-zeroes/
/*
Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Example:
Input: [0,1,0,3,12]
Output: [1,3,12,0,0]

Note:
You must do this in-place without making a copy of the array.
Minimize the total number of operations.
 */

func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}
	j := -1 // nums[0...j] != 0
	for _, num := range nums {
		if num != 0 {
			nums[j+1] = num
			j++
		}
	}
	for i := j + 1; i < len(nums); i++ {
		nums[i] = 0
	}
}

// Runtime: 84 ms, faster than 100.00% of Go online submissions for Move Zeroes.
