package main

import "strconv"

// ## Base 7 converter

// ---

// Given an integer, return its base 7 string representation.

// Free to use any programing language

// Example 1:

// ```
// Input: 20
// Output: "26"
// ```

// Example 2:

// ```
// Input: -7
// Output: "-10"
// ```
func base7trans(num int) string {
	res := ""
	nflag := ""
	if num < 0 {
		nflag = "-"
		num = -num
	}
	for {
		if num < 7 {
			res = strconv.Itoa(num) + res
			break
		}
		remain := num % 7
		res = strconv.Itoa(remain) + res
		num = (num - remain) / 7
	}
	return nflag + res
}
