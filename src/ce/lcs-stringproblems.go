/*
 * Copyright (c) 2020. Neosemantix, Inc.
 * Author: Umesh Patil
 */

package main

import (
	"strings"
	"unicode"
)

/*
 * LeetCode problem no. 125: Valid Palindrome
 * https://leetcode.com/problems/valid-palindrome/description/
 *
 * A phrase is a palindrome if, after converting all uppercase letters into lowercase letters
 * and removing all non-alphanumeric characters, it reads the same forward and backward.
 * Alphanumeric characters include letters and numbers.
 *
 * Given a string s, return true if it is a palindrome, or false otherwise.
 */
func isPalindrome(s string) bool {
	if len(s) > 0 {
		ss := removeNonAlphanumeric(strings.ToLower(s))
		runeArray := []rune(ss)
		size := len(runeArray)
		if size > 0 {
			i := 0
			midPoint := size / 2
			for i <= midPoint {
				if runeArray[i] != runeArray[size-i-1] {
					return false
				}
				i++
			}
		}
		return true // after curating the string became empty or it was palindrome when non-empty
	}
	return true
}

func removeNonAlphanumeric(input string) string {
	if len(input) > 0 {
		sb := strings.Builder{}
		for _, r := range input {
			if unicode.IsLetter(r) || unicode.IsNumber(r) {
				sb.WriteRune(r)
			}
		}
		return sb.String()
	}
	return input
}
