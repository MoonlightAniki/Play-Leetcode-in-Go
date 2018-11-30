package solution0917

// 917. Reverse Only Letters
// https://leetcode.com/problems/reverse-only-letters/
/*
Given a string S, return the "reversed" string where all characters that are not a letter stay in the same place, and all letters reverse their positions.

Example 1:
Input: "ab-cd"
Output: "dc-ba"

Example 2:
Input: "a-bC-dEf-ghIj"
Output: "j-Ih-gfE-dCba"

Example 3:
Input: "Test1ng-Leet=code-Q!"
Output: "Qedo1ct-eeLg=ntse-T!"

Note:
S.length <= 100
33 <= S[i].ASCIIcode <= 122
S doesn't contain \ or "
 */

func reverseOnlyLetters(S string) string {
	s := []rune(S)
	i, j := 0, len(S)-1
	for {
		for i < j && !isLetter(s[i]) {
			i++
		}
		for j > i && !isLetter(s[j]) {
			j--
		}
		if i >= j {
			break
		}
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	return string(s)
}

func isLetter(ch rune) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z')
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Only Letters.
