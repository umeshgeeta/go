/*
 * https://leetcode.com/problems/roman-to-integer/description/
 * Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
 *
 * Symbol       Value
 * I             1
 * V             5
 * X             10
 * L             50
 * C             100
 * D             500
 * M             1000
 *
 * For example, 2 is written as II in Roman numeral, just two ones added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.
 *
 * Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:
 *
 * I can be placed before V (5) and X (10) to make 4 and 9.
 * X can be placed before L (50) and C (100) to make 40 and 90.
 * C can be placed before D (500) and M (1000) to make 400 and 900.
 * Given a roman numeral, convert it to an integer.
 *
 * Author: Umesh Patil
 */

package main

import "fmt"

var r2imap map[rune]int

func init() {
	r2imap = make(map[rune]int)
	r2imap['I'] = 1
	r2imap['V'] = 5
	r2imap['X'] = 10
	r2imap['L'] = 50
	r2imap['C'] = 100
	r2imap['D'] = 500
	r2imap['M'] = 1000
}

func romanToInt(s string) int {
	chars := []rune(s)
	val := 0
	j := 0
	rs := len(chars)
	for j < rs-1 {
		c := chars[j]
		switch c {

		case 'I':
			if chars[j+1] == 'V' {
				val += 4
				j++
			} else if chars[j+1] == 'X' {
				val += 9
				j++
			} else {
				val += 1
			}
		case 'X':
			if chars[j+1] == 'L' {
				val += 40
				j++
			} else if chars[j+1] == 'C' {
				val += 90
				j++
			} else {
				val += 10
			}
		case 'C':
			if chars[j+1] == 'D' {
				val += 400
				j++
			} else if chars[j+1] == 'M' {
				val += 900
				j++
			} else {
				val += 100
			}
		default:
			val += r2imap[c]
		}
		j++
	}
	if j < rs {
		val += r2imap[chars[j]]
	}
	return val
}

func main() {
	fmt.Printf("%d\n", romanToInt("III"))
	fmt.Printf("%d\n", romanToInt("LVIII"))
	fmt.Printf("%d\n", romanToInt("MCMXCIV"))
	fmt.Printf("%d\n", romanToInt("MCMXCV"))

	fmt.Printf("%d\n", romanToInt("DCXXI"))
}
