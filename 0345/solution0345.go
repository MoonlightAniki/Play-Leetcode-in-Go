package _345

// 345. Reverse Vowels of a String
// https://leetcode.com/problems/reverse-vowels-of-a-string/
/*
Write a function that takes a string as input and reverse only the vowels of a string.

Example 1:
Input: "hello"
Output: "holle"

Example 2:
Input: "leetcode"
Output: "leotcede"
Note:
The vowels does not include the letter "y".
 */

func reverseVowels(s string) string {
	bytes := []byte(s)
	i, j := 0, len(bytes)-1

	for {
		for i < j && !isVowel(bytes[i]) {
			i++
		}
		for j > i && !isVowel(bytes[j]) {
			j--
		}
		if i >= j {
			break
		}
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
	return string(bytes)
}

func isVowel(ch byte) bool {
	return ch == 'a' || ch == 'A' ||
		ch == 'e' || ch == 'E' ||
		ch == 'i' || ch == 'I' ||
		ch == 'o' || ch == 'O' ||
		ch == 'u' || ch == 'U'
}

// Runtime: 4 ms, faster than 100.00% of Go online submissions for Reverse Vowels of a String.
