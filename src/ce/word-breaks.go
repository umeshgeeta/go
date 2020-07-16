// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

package main

import (
	"fmt"
	"strings"
)

func wordBreak(s string, wordDict []string) bool {
	sr := []rune(s)
	inputChars := make(map[rune]bool)
	for _, r := range sr {
		inputChars[r] = true
	}
	wdChars := make(map[rune]bool)
	for _, w := range wordDict {
		wr := []rune(w)
		for _, wrc := range wr {
			wdChars[wrc] = true
		}
	}
	for ic, _ := range inputChars {
		_, ok := wdChars[ic]
		if !ok {
			// a character in the input string is not found in any dictionary word
			return false
		}
	}
	for _, wd := range wordDict {
		if strings.HasPrefix(s, wd) {
			l := len(wd)
			if len(s) == l {
				return true
			} else {
				if wordBreak(s[l:], wordDict) {
					return true
				}
				// else need to try the next word
			}
		}
	}
	return false
}

func main() {
	input := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
	dictionary := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	fmt.Printf("%v\n", wordBreak(input, dictionary))
}
