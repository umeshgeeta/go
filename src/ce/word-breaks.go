// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

package main

import (
	"fmt"
)

func wordBreak(s string, wordDict []string) bool {
	memo := make(map[int]bool)
	wDic := make(map[string]bool)
	for _, w := range wordDict {
		wDic[w] = true
	}
	return wordBrekImpl(s, wDic, 0, memo)
}

func wordBrekImpl(s string, wDict map[string]bool, start int, memo map[int]bool) bool {
	if start == len(s) {
		return true
	}
	v, ok := memo[start]
	if ok {
		return v
	}
	ls := len(s)
	end := start + 1
	for end <= ls {
		ss := s[start:end]
		p, ok := wDict[ss]
		if ok && p && wordBrekImpl(s, wDict, end, memo) {
			memo[start] = true
			return true
		}
		end++
	}
	memo[start] = false
	return false
}

func main() {
	input := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
	dictionary := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	fmt.Printf("%v\n", wordBreak(input, dictionary))
}
