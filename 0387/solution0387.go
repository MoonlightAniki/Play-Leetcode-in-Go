package _387

// 387. First Unique Character in a String
// https://leetcode.com/problems/first-unique-character-in-a-string/
/*
Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.

Examples:
s = "leetcode"
return 0.

s = "loveleetcode",
return 2.

Note: You may assume the string contain only lowercase letters.
 */

func firstUniqChar(s string) int {
	freq := [256]int{}
	for _, ch := range s {
		freq[ch]++
	}
	for i, ch := range s {
		if freq[ch] == 1 {
			return i
		}
	}
	return -1
}

// Runtime: 8 ms, faster than 100.00% of Go online submissions for First Unique Character in a String.
