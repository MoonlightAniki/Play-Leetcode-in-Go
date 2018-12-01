package _202

// 202. Happy Number
// https://leetcode.com/problems/happy-number/
/*
Write an algorithm to determine if a number is "happy".
A happy number is a number defined by the following process: Starting with any positive integer, replace the number by
the sum of the squares of its digits, and repeat the process until the number equals 1 (where it will stay), or it loops
endlessly in a cycle which does not include 1. Those numbers for which this process ends in 1 are happy numbers.

Example:
Input: 19
Output: true
Explanation:
1^2 + 9^2 = 82
8^2 + 2^2 = 68
6^2 + 8^2 = 100
1^2 + 0^2 + 0^2 = 1
 */

func isHappy(n int) bool {
	sum := n
	rec := make(map[int]bool)
	rec[n] = true
	for {
		sum = sumSquare(sum)
		if sum == 1 {
			return true
		}
		if rec[sum] {
			return false
		} else {
			rec[sum] = true
		}
	}
}

func sumSquare(n int) int {
	sum := 0
	for n > 0 {
		r := n % 10
		n /= 10
		sum += r * r
	}
	return sum
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Happy Number.
