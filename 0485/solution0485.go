package solution0485

// 485. Max Consecutive Ones
// https://leetcode.com/problems/max-consecutive-ones/
/*
Given a binary array, find the maximum number of consecutive 1s in this array.

Example 1:
Input: [1,1,0,1,1,1]
Output: 3
Explanation: The first two digits or the last three digits are consecutive 1s.
    The maximum number of consecutive 1s is 3.

Note:
The input array will only contain 0 and 1.
The length of input array is a positive integer and will not exceed 10,000
 */
func findMaxConsecutiveOnes(nums []int) int {
	maxEndingHere := 0
	maxSoFar := 0
	for _, num := range nums {
		if num == 0 {
			maxEndingHere = 0
		} else if num == 1 {
			maxEndingHere++
			maxSoFar = max(maxSoFar, maxEndingHere)
		}
	}
	return maxSoFar
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Runtime: 56 ms, faster than 100.00% of Go online submissions for Max Consecutive Ones.
