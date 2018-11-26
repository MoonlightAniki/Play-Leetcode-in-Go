package solution0461

import "fmt"

// 461. Hamming Distance
// https://leetcode.com/problems/hamming-distance/
/*
The Hamming distance between two integers is the number of positions at which the corresponding bits are different.
Given two integers x and y, calculate the Hamming distance.

Note:
0 ≤ x, y < 231.

Example:
Input: x = 1, y = 4

Output: 2
Explanation:
1   (0 0 0 1)
4   (0 1 0 0)
       ↑   ↑

The above arrows point to positions where the corresponding bits are different.
 */

func hammingDistance(x int, y int) int {
	res := 0
	xor := x ^ y
	for xor != 0 {
		res++
		xor &= xor - 1
	}
	return res
}

func Test() {
	fmt.Println(hammingDistance(1, 4))
}

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Hamming Distance.
