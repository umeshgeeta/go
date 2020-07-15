// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

package main

import (
	"fmt"
	"strings"
)

// Given an array of integers, return indices of the two numbers such that
// they add up to a specific target. You may assume that each input would have
// exactly one solution, and you may not use the same element twice.
//
// Given nums = [2, 7, 11, 15], target = 9,
// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1].
//
// https://leetcode.com/explore/interview/card/amazon/76/array-and-strings/508/
func twoSum(nums []int, target int) []int {
	result := make([]int, 0, 2)
	m := make(map[int]int) // map of remaining value and the index
	for i, v := range nums {
		remaining := target - v
		index, ok := m[remaining]
		if ok {
			result = append(result, index, i)
			return result
		} else {
			_, present := m[v]
			if !present {
				m[v] = i
			}
		}
	}
	return result
}

func sanitize(word string) string {
	result := strings.ToLower(word)
	if strings.HasPrefix(result, "'") || strings.HasPrefix(result, "\"") {
		result = result[1:]
	}
	if strings.HasSuffix(result, ",") || strings.HasSuffix(result, ";") ||
		strings.HasSuffix(result, "?") || strings.HasSuffix(result, "!") ||
		strings.HasSuffix(result, ".") || strings.HasSuffix(result, ":") {
		l := len(result)
		result = result[0 : l-1]
	}
	if strings.HasSuffix(result, "'") || strings.HasPrefix(result, "\"") {
		l := len(result)
		result = result[0 : l-1]
	}
	return result
}

func mostCommonWord(paragraph string, banned []string) string {
	wordCount := make(map[string]int)
	words := strings.Fields(paragraph)
	for _, v := range words {
		w := sanitize(v)
		ww := strings.Split(w, ",")
		for _, vv := range ww {
			c, ok := wordCount[vv]
			if ok {
				wordCount[vv] = c + 1
			} else {
				wordCount[vv] = 1
			}
		}

	}
	for _, b := range banned {
		delete(wordCount, b)
	}
	max := 0
	maxWord := ""
	for k, v := range wordCount {
		if v > max {
			max = v
			maxWord = k
		}
	}
	fmt.Printf("%v\n", wordCount)
	return maxWord
}

func main() {

	paragraph := "L, P! X! C; u! P? w! P. G, S? l? X? D. w? m? f? v, x? i. z; x' m! U' M! j? V; l. S! j? r, K. O? k? p? p, H! t! z' X! v. u; F, h; s? X? K. y, Y! L; q! y? j, o? D' y? F' Z; E? W; W' W! n! p' U. N; w? V' y! Q; J, o! T? g? o! N' M? X? w! V. w? o' k. W. y, k; o' m! r; i, n. k, w; U? S? t; O' g' z. V. N? z, W? j! m? W! h; t! V' T! Z? R' w, w? y? y; O' w; r? q. G, V. x? n, Y; Q. s? S. G. f, s! U? l. o! i. L; Z' X! u. y, Q. q; Q, D; V. m. q. s? Y, U; p? u! q? h? O. W' y? Z! x! r. E, R, r' X' V, b. z, x! Q; y, g' j; j. q; W; v' X! J' H? i' o? n, Y. X! x? h? u; T? l! o? z. K' z' s; L? p? V' r. L? Y; V! V' S. t? Z' T' Y. s? i? Y! G? r; Y; T! h! K; M. k. U; A! V? R? C' x! X. M; z' V! w. N. T? Y' w? n, Z, Z? Y' R; V' f; V' I; t? X? Z; l? R, Q! Z. R. R, O. S! w; p' T. u? U! n, V, M. p? Q, O? q' t. B, k. u. H' T; T? S; Y! S! i? q! K' z' S! v; L. x; q; W? m? y, Z! x. y. j? N' R' I? r? V! Z; s, O? s; V, I, e? U' w! T? T! u; U! e? w? z; t! C! z? U, p' p! r. x; U! Z; u! j; T! X! N' F? n! P' t, X. s; q'"
	bn := [20]string{"m", "i", "s", "w", "y", "d", "q", "l", "a", "p", "n", "t", "u", "b", "o", "e", "f", "g", "c", "x"}
	banned := make([]string, 0, 30)
	for _, v := range bn {
		banned = append(banned, v)
	}
	fmt.Printf("answer: %v\n", mostCommonWord(paragraph, banned))

	nums := [4]int{2, 7, 11, 15}
	fmt.Printf("answer: %v\n", twoSum(nums[0:4], 9))
	nums = [4]int{2, 5, 5, 11}
	fmt.Printf("answer: %v\n", twoSum(nums[0:4], 10))
}
