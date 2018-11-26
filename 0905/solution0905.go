package solution0905

import "fmt"

// 905. Sort Array By Parity
// https://leetcode.com/problems/sort-array-by-parity/
/*
Given an array A of non-negative integers, return an array consisting of all the even elements of A, followed by all the odd elements of A.
You may return any answer array that satisfies this condition.

Example 1:
Input: [3,1,2,4]
Output: [2,4,3,1]
The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.

Note:
1 <= A.length <= 5000
0 <= A[i] <= 5000
 */

func sortArrayByParity(A []int) []int {
	even := 0
	odd := len(A) - 1
	for {
		for even < len(A) && A[even]%2 == 0 {
			even++
		}
		for odd >= 0 && A[odd]%2 == 1 {
			odd--
		}
		if even > odd {
			break
		}
		A[even], A[odd] = A[odd], A[even]
	}
	return A
}

func Test() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(sortArrayByParity(nums))
}

// Runtime: 72 ms, faster than 100.00% of Go online submissions for Sort Array By Parity.
