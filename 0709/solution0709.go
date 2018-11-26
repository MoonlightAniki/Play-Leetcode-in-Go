package solution0709

// 709. To Lower Case
// https://leetcode.com/problems/to-lower-case/
/*
mplement function ToLowerCase() that has a string parameter str, and returns the same string in lowercase.

Example 1:
Input: "Hello"
Output: "hello"

Example 2:
Input: "here"
Output: "here"

Example 3:
Input: "LOVELY"
Output: "lovely"
 */

func toLowerCase2(str string) string {
	if len(str) == 0 {
		return str
	}
	for i := range str {
		if isUpper(str[i]) {

		}
	}
	return str
}

func isUpper(ch uint8) bool {
	return ch >= 'A' && ch <= 'Z'
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for To Lower Case.
