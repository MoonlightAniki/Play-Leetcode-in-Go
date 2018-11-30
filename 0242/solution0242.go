package _242

import "reflect"

// 242. Valid Anagram
// https://leetcode.com/problems/valid-anagram/
/*
Given two strings s and t , write a function to determine if t is an anagram of s.

Example 1:
Input: s = "anagram", t = "nagaram"
Output: true

Example 2:
Input: s = "rat", t = "car"
Output: false

Note:
You may assume the string contains only lowercase alphabets.

Follow up:
What if the inputs contain unicode characters? How would you adapt your solution to such case?
 */

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	return reflect.DeepEqual(getFreq(s), getFreq(t))
}

func getFreq(s string) map[rune]int {
	freq := make(map[rune]int, 26)
	for _, ch := range s {
		freq[ch]++
	}
	return freq
}

// Runtime: 4 ms, faster than 73.64% of Go online submissions for Valid Anagram.
