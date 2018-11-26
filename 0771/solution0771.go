package solution0771

import "fmt"

// 771. Jewels and Stones
// https://leetcode.com/problems/jewels-and-stones/
/*
You're given strings J representing the types of stones that are jewels, and S representing the stones you have.  Each
character in S is a type of stone you have.  You want to know how many of the stones you have are also jewels.
The letters in J are guaranteed distinct, and all characters in J and S are letters. Letters are case sensitive,
so "a" is considered a different type of stone from "A".

Example 1:
Input: J = "aA", S = "aAAbbbb"
Output: 3

Example 2:
Input: J = "z", S = "ZZ"
Output: 0

Note:
S and J will consist of letters and have length at most 50.
The characters in J are distinct.
 */
func numJewelsInStones(J string, S string) int {
	if len(J) == 0 || len(S) == 0 {
		return 0
	}
	freq := make([]int, 256)
	for _, ch := range J {
		freq[ch]++
	}
	res := 0
	for _, ch := range S {
		if freq[ch] != 0 {
			res++
		}
	}
	return res
}

func Test() {
	J := "aA"
	S := "aAAbbbb"
	fmt.Println(numJewelsInStones(J, S))
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Jewels and Stones.
