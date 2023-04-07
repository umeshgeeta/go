/*
 * Copyright (c) 2020. Neosemantix, Inc.
 * Author: Umesh Patil
 */

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello World"
	fmt.Printf("%d\n", lengthOfLastWord(s))

	s = "   fly me   to   the moon  "
	fmt.Printf("%d\n", lengthOfLastWord(s))

	s = "luffy is still joyboy"
	fmt.Printf("%d\n", lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {
	words := strings.Fields(s)
	lastWordLength := 0
	wordCount := len(words)
	if wordCount > 0 {
		lastWordLength = len(words[wordCount-1])
	}
	return lastWordLength
}
