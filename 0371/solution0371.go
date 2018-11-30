package _371

// 371. Sum of Two Integers
// https://leetcode.com/problems/sum-of-two-integers/
/*
Calculate the sum of two integers a and b, but you are not allowed to use the operator + and -.

Example 1:
Input: a = 1, b = 2
Output: 3

Example 2:
Input: a = -2, b = 3
Output: 1
 */

func getSum(a int, b int) int {
	for a != 0 {
		a, b = (a&b)<<1, a^b
	}
	return b
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Sum of Two Integers.
