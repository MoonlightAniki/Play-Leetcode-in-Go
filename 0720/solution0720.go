package _720

import (
	"sort"
)

// 720. Longest Word in Dictionary
// https://leetcode.com/problems/longest-word-in-dictionary/
/*
Given a list of strings words representing an English Dictionary, find the longest word in words that can be built one
character at a time by other words in words. If there is more than one possible answer, return the longest word with the smallest lexicographical order.
If there is no answer, return the empty string.

Example 1:
Input:
words = ["w","wo","wor","worl", "world"]
Output: "world"
Explanation:
The word "world" can be built one character at a time by "w", "wo", "wor", and "worl".

Example 2:
Input:
words = ["a", "banana", "app", "appl", "ap", "apply", "apple"]
Output: "apple"
Explanation:
Both "apply" and "apple" can be built from other words in the dictionary. However, "apple" is lexicographically smaller than "apply".
 */

func longestWord(words []string) string {
	sort.Strings(words)
	//fmt.Println(words)
	records := make(map[string]bool)
	res := ""
	records[res] = true
	for _, w := range words {
		if records[w[0:len(w)-1]] {
			records[w] = true
			if len(w) > len(res) {
				res = w
			}
		}
	}
	return res
}

// Runtime: 20 ms, faster than 94.44% of Go online submissions for Longest Word in Dictionary.
