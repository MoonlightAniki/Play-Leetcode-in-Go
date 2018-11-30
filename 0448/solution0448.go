package solution0448

// 448. Find All Numbers Disappeared in an Array
// https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/
/*
Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.
Find all the elements of [1, n] inclusive that do not appear in this array.
Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.

Example:
Input:
[4,3,2,7,8,2,3,1]
Output:
[5,6]
 */
func findDisappearedNumbers(nums []int) []int {
	for _, num := range nums {
		if nums[abs(num)-1] > 0 {
			nums[abs(num)-1] *= -1
		}
	}
	res := make([]int, 0)
	for i, num := range nums {
		if num > 0 {
			res = append(res, i+1)
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Runtime: 496 ms, faster than 84.81% of Go online submissions for Find All Numbers Disappeared in an Array.