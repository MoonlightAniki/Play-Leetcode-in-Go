package solution0557

import (
	"strings"
	"fmt"
)

// 557. Reverse Words in a String III
// https://leetcode.com/problems/reverse-words-in-a-string-iii/
/*
Given a string, you need to reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.

Example 1:
Input: "Let's take LeetCode contest"
Output: "s'teL ekat edoCteeL tsetnoc"

Note: In the string, each word is separated by single space and there will not be any extra space in the string.
 */
func reverseWords(s string) string {
	words := strings.Split(s, " ")
	for i := range words {
		words[i] = reverse(words[i])
	}
	return strings.Join(words, " ")
}

func reverse(s string) string {
	runes := []rune(s)
	i, j := 0, len(runes)-1
	for i < j {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}
	return string(runes)
}

func Test() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}

// Runtime: 12 ms, faster than 63.08% of Go online submissions for Reverse Words in a String III.
