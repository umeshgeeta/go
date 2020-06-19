// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

// Leetcode problem: 	https://leetcode.com/problems/goat-latin/
package main

import (
	"fmt"
	"strings"
)

func toGoatLatin(S string) string {
	words := strings.Fields(S)
	if len(words) > 0 {
		result := ""
		for i, aWord := range words {
			result = result + goatLatinWord(aWord, i+1) + " "
		}
		return strings.Trim(result, " ")
	} else {
		return S
	}
}

func goatLatinWord(word string, whereInSentence int) string {
	result := ""
	var firstChar rune = rune(word[0])
	if isVowel(firstChar) {
		result = word + "ma"
	} else {
		result = word[1:]
		result = result + string(firstChar) + "ma"
	}
	for i := 0; i < whereInSentence; i++ {
		result = result + string('a')
	}
	return result
}

func isVowel(r rune) bool {
	if r == 'a' || r == 'A' || r == 'e' || r == 'E' || r == 'i' || r == 'I' || r == 'o' || r == 'O' || r == 'u' || r == 'U' {
		return true
	} else {
		return false
	}
}

func main() {
	input := "I speak Goat Latin"
	fmt.Println(toGoatLatin(input))

	input = "The quick brown fox jumped over the lazy dog"
	fmt.Println(toGoatLatin(input))
}
