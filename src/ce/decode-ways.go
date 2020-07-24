// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

package main

import "fmt"

// A message containing letters from A-Z is being encoded to numbers using the
// following mapping:
//
// 'A' -> 1
// 'B' -> 2
// ...
// 'Z' -> 26
//
// Given a non-empty string containing only digits, determine the total number
// of ways to decode it.
//
// https://leetcode.com/explore/interview/card/facebook/55/dynamic-programming-3/264/

const r0 rune = '0'
const r1 rune = '1'
const r2 rune = '2'
const r7 rune = '7'

func numDecodings(s string) int {
	result := 0
	l := len(s)
	switch l {
	case 1:
		var r rune = rune(s[0])
		if r > r0 {
			result = 1
		}
		// else it is zero
	case 2:
		var cs []rune = []rune(s)
		d1 := cs[0]
		d2 := cs[1]
		switch d1 {
		case r0:
			// we return with 0 value
		case r1:
			if d2 != r0 {
				result = 2
			} else {
				result = 1
			}
		case r2:
			if d2 < r7 && d2 > r0 {
				result = 2
			} else {
				result = 1
			}
		default:
			if d2 != r0 {
				result = 1
			}
		}
	default:
		{
			var cs []rune = []rune(s)
			d1 := cs[0]
			d2 := cs[1]
			switch d1 {
			case r0:
				//result = numDecodings(s[1:])
				// we return with 0 value
			case r1:
				result = numDecodings(s[1:])
				result = result + numDecodings(s[2:])
			case r2:
				result = numDecodings(s[1:])
				if d2 < r7 {
					result = result + numDecodings(s[2:])
				}
			default:
				if d2 != r0 {
					result = numDecodings(s[1:])
				}
				// else first digit is 3 or more and the second one is 0; we return 0 with no recursion
			}
		}
	}
	return result
}

func main() {
	s := "12"
	fmt.Printf("%d\n", numDecodings(s))
	s = "611"
	fmt.Printf("%d\n", numDecodings(s))
	s = "301"
	fmt.Printf("%d\n", numDecodings(s))
	s = "230"
	fmt.Printf("%d\n", numDecodings(s))
	s = "01"
	fmt.Printf("%d\n", numDecodings(s))
	s = "10"
	fmt.Printf("%d\n", numDecodings(s))
	s = "226"
	fmt.Printf("%d\n", numDecodings(s))
	s = "21212"
	fmt.Printf("%d\n", numDecodings(s))
}
