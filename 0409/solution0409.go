package _409

// 409. Longest Palindrome
// https://leetcode.com/problems/longest-palindrome/
/*
Given a string which consists of lowercase or uppercase letters, find the length of the longest palindromes that can be built with those letters.
This is case sensitive, for example "Aa" is not considered a palindrome here.

Note:
Assume the length of given string will not exceed 1,010.

Example:
Input:
"abccccdd"
Output:
7
Explanation:
One longest palindrome that can be built is "dccaccd", whose length is 7.
 */

func longestPalindrome(s string) int {
	freq := [256]int{}
	for _, ch := range s {
		freq[ch]++
	}
	res := 0
	oddCount := 0
	for _, count := range freq {
		if count%2 == 0 {
			res += count
		} else {
			res += count - 1
			oddCount++
		}
	}
	if oddCount > 0 {
		res += 1
	}
	return res
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Longest Palindrome.
