package solution0412

import "strconv"

// 412. Fizz Buzz
// https://leetcode.com/problems/fizz-buzz/
/*
Write a program that outputs the string representation of numbers from 1 to n.
But for multiples of three it should output “Fizz” instead of the number and for the multiples of five output “Buzz”. For numbers which are multiples of both three and five output “FizzBuzz”.

Example:
n = 15,
Return:
[
    "1",
    "2",
    "Fizz",
    "4",
    "Buzz",
    "Fizz",
    "7",
    "8",
    "Fizz",
    "Buzz",
    "11",
    "Fizz",
    "13",
    "14",
    "FizzBuzz"
]
 */
func fizzBuzz(n int) []string {
	res := make([]string, 0)
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			res = append(res, "FizzBuzz")
		} else if i%5 == 0 {
			res = append(res, "Buzz")
		} else if i%3 == 0 {
			res = append(res, "Fizz")
		} else {
			res = append(res, strconv.Itoa(i))
		}
	}
	return res

/*	res := make([]string, n)
	for i := range res {
		x := i + 1
		if x%15 == 0 {
			res[i] = "FizzBuzz"
		} else if x%5 == 0 {
			res[i] = "Buzz"
		} else if x%3 == 0 {
			res[i] = "Fizz"
		} else {
			res[i] = strconv.Itoa(x)
		}
	}
	return res*/
}
// Runtime: 140 ms, faster than 100.00% of Go online submissions for Fizz Buzz.
