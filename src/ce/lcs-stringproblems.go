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
