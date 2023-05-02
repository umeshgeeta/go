/*
 * Copyright (c) 2020. Neosemantix, Inc.
 * Author: Umesh Patil
 */

package main

import (
	"fmt"
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

/*
 * LeetCode problem no. 2423: Remove letter to equalize frequency
 * https://leetcode.com/problems/remove-letter-to-equalize-frequency/description/
 *
 * You are given a 0-indexed string word, consisting of lowercase English letters.
 * You need to select one index and remove the letter at that index from word so that
 * the frequency of every letter present in word is equal.
 *
 * Return true if it is possible to remove one letter so that the frequency of all letters
 * in word are equal, and false otherwise.
 *
 * (Wickedly complex, devilish problem!)
 */
func equalFrequency(word string) bool {
	size := len(word)
	if size > 0 {
		letters := make(map[rune]int)
		ra := []rune(word)
		maxfreq := 0
		for _, r := range ra {
			c, ok := letters[r]
			if ok {
				letters[r] = c + 1
			} else {
				letters[r] = 1
			}
			if maxfreq < letters[r] {
				maxfreq = letters[r]
			}
		}
		if maxfreq == 1 {
			// all letters of frequency 1
			// removing any one letter still leave remaining letters
			// with frequency 1; we should be good here
			return true
		}
		// also if there is only one letter, we are good too
		if len(letters) == 1 {
			return true
		}
		frequencyCount := make(map[int]int)
		minfreqVal := size
		maxfreqVal := 0
		for _, f := range letters {
			fc, ok := frequencyCount[f]
			if ok {
				frequencyCount[f] = fc + 1
			} else {
				frequencyCount[f] = 1
			}
			if f < minfreqVal {
				minfreqVal = f
			}
			if f > maxfreqVal {
				maxfreqVal = f
			}
		}
		howManyFrequencies := len(frequencyCount)
		switch howManyFrequencies {

		case 1:
			// all letters of the same frequency more then 1
			// cannot remove a single letter to make them equal
			return false

		case 2:
			howManyminfreqCount := frequencyCount[minfreqVal]
			if minfreqVal == 1 {
				if howManyminfreqCount == 1 {
					// there it only one letter of frequency 1
					// that can be removed and all other remaining letters have same frequency
					return true
				}
				// else there are multiple letters with one frequency with at least one of diff frequency
				// but
				return false
			}
			// minimum frequency is not 1
			// so we need letter with more than minimum frequency exactly 1 AND
			// difference between min & max to be exactly 1 (since min > 1)
			if maxfreqVal == minfreqVal+1 && frequencyCount[maxfreqVal] == 1 {
				return true
			}
			return false

		default:
			// there are more than 2 frequencies
			// there cannot be a single letter removal to make them same
			return false
		}
	}
	return false
}

/*
 * LeetCode problem no. 3: Longest Substring Without Repeating Characters
 *
 * Given a string s, find the length of the longest substring without repeating characters.
 */
func lengthOfLongestSubstring(s string) int {
	size := len(s)
	if size < 2 {
		return size
	}
	charsScannedSoFar := make(map[rune]int)
	i := 0
	lols := 0 // length of the longest substring
	start := 0
	sr := []rune(s)
	for i < size {
		j, ok := charsScannedSoFar[sr[i]]
		if !ok {
			charsScannedSoFar[sr[i]] = i
			lols = checkLength(charsScannedSoFar, lols)
		} else {
			// we came across a repeat char
			firstpart := j - start + 1
			if firstpart > lols {
				lols = firstpart
			}
			// new substring will start from next char of the earlier
			start = j + 1
			// we need to remove all chars with index upto j
			for in, ch := range charsScannedSoFar {
				if ch <= j {
					delete(charsScannedSoFar, in)
				}
			}
			// now we add the current char
			charsScannedSoFar[sr[i]] = i

			remainingChars := len(charsScannedSoFar)
			if lols < remainingChars {
				lols = remainingChars
			}
			lols = checkLength(charsScannedSoFar, lols)
		}
		i++
	}
	lols = checkLength(charsScannedSoFar, lols)
	return lols
}

func checkLength(charsScannedSoFar map[rune]int, lols int) int {
	remainingChars := len(charsScannedSoFar)
	if lols < remainingChars {
		lols = remainingChars
	}
	return lols
}

// LeetCode no. 926: Flip String to Monotone Increasing
// https://leetcode.com/problems/flip-string-to-monotone-increasing/description/
// A binary string is monotone increasing if it consists of some number of 0's (possibly none),
// followed by some number of 1's (also possibly none).
//
// You are given a binary string s. You can flip s[i] changing it from 0 to 1 or from 1 to 0.
//
// Return the minimum number of flips to make s monotone increasing.
// Constraints:
//
// 1 <= s.length <= 105
// s[i] is either '0' or '1'.
func minFlipsMonoIncr(s string) int {
	zero2one := howManyZeros(s)
	one2zero := 0
	minCount := zero2one + one2zero // which is same as zero2one
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			zero2one--
		} else {
			one2zero++
		}
		if minCount > zero2one+one2zero {
			minCount = zero2one + one2zero
		}
	}
	return minCount
}

func howManyZeros(s string) int {
	zc := 0
	for _, c := range s {
		if c == '0' {
			zc++
		}
	}
	return zc
}

// LeetCode problem no. 22: Generate Parentheses (not clear why LeetCode gives an issue for this impl.)
// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
// Constraints: 1 <= n <= 8
func generateParenthesisVer1(n int) []string {
	result := make([]string, 0)
	if n == 1 {
		result = append(result, "()")
		return result
	}
	previous := generateParenthesisVer1(n - 1)
	for _, p := range previous {
		result = append(result, "("+p+")")
		left := "()" + p
		right := p + "()"
		result = append(result, left)
		if !strings.EqualFold(left, right) {
			result = append(result, right)
		}
	}
	return result
}

// LeetCode problem no. 22: Generate Parentheses (submitted to LeedCode)
// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
// Constraints: 1 <= n <= 8
func generateParenthesis(n int) []string {
	result := []string{}
	if n == 0 {
		result = append(result, "")
		return result
	}
	for lc := 0; lc < n; lc++ {
		for _, sl := range generateParenthesis(lc) {
			for _, sr := range generateParenthesis(n - 1 - lc) {
				result = append(result, "("+sl+")"+sr)
			}
		}
	}
	return result
}

func main() {
	//fmt.Println(lengthOfLongestSubstring("abcb"))
	fmt.Println(minFlipsMonoIncr("00110"))
	fmt.Println(minFlipsMonoIncr("010110"))
	fmt.Println(minFlipsMonoIncr("00011000"))
}
